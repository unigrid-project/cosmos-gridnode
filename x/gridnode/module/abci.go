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
	storeAdapter := runtime.KVStoreAdapter(k.GetStoreService().OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	iterator := storetypes.KVStorePrefixIterator(store, []byte(k.GetBondingPrefix()))

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

		var entriesChanged bool

		// Filter out entries that have completed unbonding
		newEntries := make([]types.UnbondingEntry, 0, len(entries))
		for _, entry := range entries {
			timestamp := time.Unix(entry.CompletionTime, 0)

			if currentTime.After(timestamp) {
				fmt.Printf("Processing unbonding for delegator: %s, amount: %d\n", entry.Account, entry.Amount)
				// Fetch the public key along with the current locked balance

				bankKeeper := k.GetBankKeeper()
				// Process the unbonding
				delegatorAddr, err := sdk.AccAddressFromBech32(entry.Account)
				if err != nil {
					fmt.Printf("Error processing unbonding for delegator %s: %v\n", entry.Account, err)
					continue
				}

				amount := math.NewInt(entry.Amount)
				coin := sdk.NewCoin("uugd", amount)
				snd := bankKeeper.SendCoinsFromModuleToAccount(goCtx, types.ModuleName, delegatorAddr, sdk.NewCoins(coin))
				if snd != nil {
					fmt.Println("Error sending coins from module to account:", err)
					continue
				}
				// Get the current delegated and locked amounts
				currentDelegatedAmount := k.GetDelegatedAmount(goCtx, delegatorAddr)
				currentLockedBalance := k.GetLockedBalance(goCtx, delegatorAddr)

				// Calculate new amounts after unbonding
				newDelegatedAmount := currentDelegatedAmount.Sub(amount)
				newLockedBalance := currentLockedBalance.Sub(amount)

				// Get the public key for the delegator
				accountPublicKey, errPk := k.GetPublicKeyForDelegator(goCtx, delegatorAddr)
				if errPk != nil {
					fmt.Printf("Error retrieving public key for delegator %s: %v\n", entry.Account, errPk)
					continue
				}
				fmt.Printf("Public key for delegator %s: %s\n", entry.Account, accountPublicKey)

				// Update the stored delegated amount
				k.SetDelegatedAmount(goCtx, delegatorAddr, newDelegatedAmount)

				// Update the stored locked balance with the public key
				k.SetLockedBalance(goCtx, delegatorAddr, newLockedBalance, accountPublicKey)

				// Log for confirmation
				fmt.Printf("Updated balance and delegation for delegator %s\n", entry.Account)

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
				entriesChanged = true
			} else {
				newEntries = append(newEntries, entry)
			}
		}

		// Update the store with the new list of unbonding entries
		if entriesChanged {
			if len(newEntries) == 0 {
				fmt.Printf("All unbonding entries processed for key: %s. Deleting key from store.\n", key)
				store.Delete(key)
			} else {
				newBz, err := json.Marshal(newEntries)
				if err != nil {
					fmt.Printf("Error marshalling new entries for key %x: %v\n", key, err)
					continue
				}

				// Log the data that is about to be stored
				fmt.Printf("Updating store for key %x with data: %s\n", key, string(newBz))

				store.Set(key, newBz)
				fmt.Printf("Updated store for key %x\n", key)
			}
		}
	}

	//fmt.Println("BeginBlocker gridnode completed.")
}
