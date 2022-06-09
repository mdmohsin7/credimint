package keeper

import (
	"encoding/binary"

	"credimint/x/credimint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetUserCredibilityCount get the total number of userCredibility
func (k Keeper) GetUserCredibilityCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserCredibilityCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetUserCredibilityCount set the total number of userCredibility
func (k Keeper) SetUserCredibilityCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.UserCredibilityCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendUserCredibility appends a userCredibility in the store with a new id and update the count
func (k Keeper) AppendUserCredibility(
	ctx sdk.Context,
	userCredibility types.UserCredibility,
) uint64 {
	// Create the userCredibility
	count := k.GetUserCredibilityCount(ctx)

	// Set the ID of the appended value
	userCredibility.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCredibilityKey))
	appendedValue := k.cdc.MustMarshal(&userCredibility)
	store.Set(GetUserCredibilityIDBytes(userCredibility.Id), appendedValue)

	// Update userCredibility count
	k.SetUserCredibilityCount(ctx, count+1)

	return count
}

// SetUserCredibility set a specific userCredibility in the store
func (k Keeper) SetUserCredibility(ctx sdk.Context, userCredibility types.UserCredibility) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCredibilityKey))
	b := k.cdc.MustMarshal(&userCredibility)
	store.Set(GetUserCredibilityIDBytes(userCredibility.Id), b)
}

// GetUserCredibility returns a userCredibility from its id
func (k Keeper) GetUserCredibility(ctx sdk.Context, id uint64) (val types.UserCredibility, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCredibilityKey))
	b := store.Get(GetUserCredibilityIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserCredibility removes a userCredibility from the store
func (k Keeper) RemoveUserCredibility(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCredibilityKey))
	store.Delete(GetUserCredibilityIDBytes(id))
}

// GetAllUserCredibility returns all userCredibility
func (k Keeper) GetAllUserCredibility(ctx sdk.Context) (list []types.UserCredibility) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCredibilityKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserCredibility
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetUserCredibilityIDBytes returns the byte representation of the ID
func GetUserCredibilityIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUserCredibilityIDFromBytes returns ID in uint64 format from a byte array
func GetUserCredibilityIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
