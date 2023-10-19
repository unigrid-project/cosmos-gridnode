package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

type KeyInfo struct {
	Keys      []string
	Timestamp int64
	UniqueId  string
	Status    string
}

func (k *Keeper) RegisterKeys(ctx sdk.Context, account sdk.AccAddress, keys []string, timestamp int64, uniqueId string) error {
	keyInfo := types.KeyInfo{
		Keys:      keys,
		Timestamp: timestamp,
		UniqueId:  uniqueId,
		Status:    "active",
	}
	return k.SetAccountKeys(ctx, account, keyInfo) // updated the function signature
}

func (k Keeper) SetAccountKeys(ctx sdk.Context, account sdk.AccAddress, info types.KeyInfo) error {
	store := ctx.KVStore(k.storeKey)
	key := k.keyForAccount(account)
	bz, err := json.Marshal(&info)
	if err != nil {
		k.Logger(ctx).Error("failed to marshal KeyInfo", "error", err)
		return err
	}
	store.Set(key, bz)
	return nil
}

func (k Keeper) GetAccountKeys(ctx sdk.Context, account sdk.AccAddress) (types.KeyInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := k.keyForAccount(account)
	bz := store.Get(key)
	if bz == nil {
		return types.KeyInfo{}, false
	}
	var info types.KeyInfo
	err := json.Unmarshal(bz, &info)
	if err != nil {
		// You might want to log the error or handle it in a way that's appropriate for your use case
		return types.KeyInfo{}, false
	}
	return info, true
}

func (k Keeper) UpdateKeyStatus(ctx sdk.Context, account sdk.AccAddress, newStatus string) error {
	info, found := k.GetAccountKeys(ctx, account)
	if !found {
		return fmt.Errorf("no keys found for account %s", account)
	}
	info.Status = newStatus
	return k.SetAccountKeys(ctx, account, info)
}

func (k Keeper) InvalidateKeys(ctx sdk.Context, account sdk.AccAddress) error {
	return k.UpdateKeyStatus(ctx, account, "invalid")
}

func (k Keeper) keyForAccount(account sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("account-keys-%s", account.String()))
}

func (k Keeper) AllKeys(ctx context.Context, req *types.QueryAllKeysRequest) (*types.QueryAllKeysResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, nil) // assuming all keys are relevant; adjust the prefix as necessary

	var accountKeysEntries []types.AccountKeysEntry

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var keyInfo types.KeyInfo
		err := json.Unmarshal(iterator.Value(), &keyInfo)
		if err != nil {
			return nil, err // or however you want to handle errors
		}

		account := sdk.AccAddress(iterator.Key()) // assuming the key is the account address; adjust as necessary
		accountKeysEntry := types.AccountKeysEntry{
			AccountAddress: account.String(),
			Keys:           keyInfo.Keys,
		}
		accountKeysEntries = append(accountKeysEntries, accountKeysEntry)
	}

	// Convert the accountKeysEntries slice to a slice of pointers to AccountKeysEntry
	var accountKeysEntryPointers []*types.AccountKeysEntry
	for i := range accountKeysEntries {
		accountKeysEntryPointers = append(accountKeysEntryPointers, &accountKeysEntries[i])
	}

	return &types.QueryAllKeysResponse{
		AccountKeysEntries: accountKeysEntryPointers,
	}, nil
}

// AccountKeys method updated to accept context.Context
func (k Keeper) AccountKeys(ctx context.Context, req *types.QueryAccountKeysRequest) (*types.QueryAccountKeysResponse, error) {
	// Convert context.Context to sdk.Context
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	account, err := sdk.AccAddressFromBech32(req.AccountAddress)
	if err != nil {
		return nil, err
	}

	info, found := k.GetAccountKeys(sdkCtx, account) // Pass sdkCtx instead of ctx
	if !found {
		return nil, fmt.Errorf("no keys found for account %s", req.AccountAddress)
	}

	return &types.QueryAccountKeysResponse{
		Keys: info.Keys,
	}, nil
}
