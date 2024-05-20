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
	currentTime := ctx.BlockTime()

	// Debugging information
	fmt.Println("BeginBlocker started")

	storeService := k.GetStoreService()
	if storeService == nil {
		panic("storeService is nil")
	}

	storeAdapter := runtime.KVStoreAdapter(k.GetStoreService().OpenKVStore(goCtx))
	if storeAdapter == nil {
		panic("storeAdapter is nil")
	}

	store := prefix.NewStore(storeAdapter, []byte(types.StoreKey))

	iterator := storetypes.KVStorePrefixIterator(store, []byte(k.GetBondingPrefix()))
	defer iterator.Close()

	// Debugging the iterator
	if iterator == nil {
		panic("iterator is nil")
	}

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		bz := iterator.Value()

		fmt.Printf("Processing key: %x, value (raw): %x\n", key, bz)

		var entries []types.UnbondingEntry
		if err := json.Unmarshal(bz, &entries); err != nil {
			fmt.Printf("Error unmarshalling unbonding entry for key %x: %v\n", key, err)
			continue
		}

		var entriesChanged bool
		newEntries := make([]types.UnbondingEntry, 0, len(entries))
		for _, entry := range entries {
			timestamp := time.Unix(entry.CompletionTime, 0)

			if currentTime.After(timestamp) {
				fmt.Printf("Processing unbonding for delegator: %s, amount: %d\n", entry.Account, entry.Amount)

				bankKeeper := k.GetBankKeeper()
				if bankKeeper == nil {
					fmt.Println("bankKeeper is nil")
					continue
				}

				delegatorAddr, err := sdk.AccAddressFromBech32(entry.Account)
				if err != nil {
					fmt.Printf("Error processing unbonding for delegator %s: %v\n", entry.Account, err)
					continue
				}

				amount := math.NewInt(entry.Amount)
				coin := sdk.NewCoin("uugd", amount)
				snd := bankKeeper.SendCoinsFromModuleToAccount(goCtx, types.ModuleName, delegatorAddr, sdk.NewCoins(coin))
				if snd != nil {
					fmt.Println("Error sending coins from module to account:", snd)
					continue
				}

				currentDelegatedAmount := k.GetDelegatedAmount(goCtx, delegatorAddr)
				currentLockedBalance := k.GetLockedBalance(goCtx, delegatorAddr)

				if currentDelegatedAmount.IsNil() || currentLockedBalance.IsNil() {
					fmt.Println("Current delegated amount or locked balance is nil")
					continue
				}

				newDelegatedAmount := currentDelegatedAmount.Sub(amount)
				newLockedBalance := currentLockedBalance.Sub(amount)

				accountPublicKey, errPk := k.GetPublicKeyForDelegator(goCtx, delegatorAddr)
				if errPk != nil {
					fmt.Printf("Error retrieving public key for delegator %s: %v\n", entry.Account, errPk)
					continue
				}

				if len(accountPublicKey) == 0 {
					fmt.Println("Account public key is empty")
					continue
				}

				k.SetDelegatedAmount(goCtx, delegatorAddr, newDelegatedAmount)
				k.SetLockedBalance(goCtx, delegatorAddr, newLockedBalance, accountPublicKey)
				fmt.Printf("Updated balance and delegation for delegator %s\n", entry.Account)
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

				fmt.Printf("Updating store for key %x with data: %s\n", key, string(newBz))
				store.Set(key, newBz)
				fmt.Printf("Updated store for key %x\n", key)
			}
		}
	}

	fmt.Println("BeginBlocker gridnode completed.")
}
