package gridnode

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func BeginBlocker(goCtx context.Context, k keeper.Keeper) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// disable heartbeat manager for now
	//k.StartHeartbeatTimer(ctx)
	// Get the current block time
	currentTime := ctx.BlockTime()

	//fmt.Println("BeginBlocker started. Current block time:", currentTime)
	//fmt.Println("CTX BlockHeight:", ctx.BlockHeight())
	// Iterate over all unbonding entries
	storeAdapter := runtime.KVStoreAdapter(k.GetStoreService().OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	iterator := storetypes.KVStorePrefixIterator(store, []byte(k.GetBondingPrefix()))
	fmt.Println("After store error:", store)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		bz := iterator.Value()

		// Log the key and the raw value
		fmt.Printf("Processing key: %x, value (raw): %x\n", key, bz)

		var entries []types.UnbondingEntry
		if err := json.Unmarshal(bz, &entries); err != nil {
			fmt.Printf("Error unmarshalling unbonding entry for key %x: %v\n", key, err)
			continue
		}

		// Filter out entries that have completed unbonding
		newEntries := make([]types.UnbondingEntry, 0, len(entries))
		for _, entry := range entries {
			timestamp := time.Unix(entry.CompletionTime, 0)

			if currentTime.After(timestamp) {
				//fmt.Printf("Processing unbonding for delegator: %s, amount: %s\n", entry.Account, entry.Amount)
				bankKeeper := k.GetBankKeeper()
				// Process the unbonding
				delegatorAddr, err := sdk.AccAddressFromBech32(entry.Account)
				if err != nil {
					fmt.Printf("Error processing unbonding for delegator %s: %v\n", entry.Account, err)
					continue
				}
				amount := math.NewInt(entry.Amount)
				coin := sdk.NewCoin("uugd", amount)
				snd := bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, delegatorAddr, sdk.NewCoins(coin))
				if snd != nil {
					fmt.Println("Error sending coins from module to account:", err)
					continue
				}
				// Reduce the delegated amount from the store
				currentDelegatedAmount := k.GetDelegatedAmount(ctx, delegatorAddr)
				newDelegatedAmount := currentDelegatedAmount.Sub(amount)
				k.SetDelegatedAmount(ctx, delegatorAddr, newDelegatedAmount)

				// Placeholder to call hedgehog
				fmt.Printf("Placeholder: Notify hedgehog that account %s is unbonding %d tokens.\n", entry.Account, entry.Amount)
				// TODO: Implement the actual call to hedgehog here

				// Emit an event for successful unbonding
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.EventTypeCompleteUnbond,
						sdk.NewAttribute(types.AttributeKeyDelegator, entry.Account),
						sdk.NewAttribute(types.AttributeKeyAmount, strconv.FormatInt(entry.Amount, 10)),
					),
				)
			} else {
				newEntries = append(newEntries, entry)
			}
		}

		// Update the store with the new list of unbonding entries
		if len(newEntries) == 0 {
			fmt.Printf("All unbonding entries processed for key: %s. Deleting key from store.\n", key)
			store.Delete(key)
		} else {
			newBz, err := json.Marshal(newEntries)
			if err != nil {
				fmt.Printf("Error marshalling new entries for key %x: %v\n", key, err)
				continue
			}
			store.Set(key, newBz)
			fmt.Printf("Updated store for key %x\n", key)
		}
	}

	fmt.Println("BeginBlocker gridnode completed.")
}
