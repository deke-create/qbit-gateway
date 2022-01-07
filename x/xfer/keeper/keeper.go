package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		typeStoreKey sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute

	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	typeStoreKey sdk.StoreKey,

	// this line is used by starport scaffolding # ibc/keeper/parameter

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		typeStoreKey: typeStoreKey,

		// this line is used by starport scaffolding # ibc/keeper/return

	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}


// SetState sets the state of a contract
func (k Keeper) SetState(ctx sdk.Context, id []byte, data []byte, typeName []byte) {

	store := ctx.KVStore(k.storeKey)

	store.Set(id, data)

	typeStore := ctx.KVStore(k.typeStoreKey)
	typeStore.Set(id, typeName)

}

// GetState sets the state of a contract
func (k Keeper) GetState(ctx sdk.Context, id []byte) []byte {
	store := ctx.KVStore(k.storeKey)
	return store.Get(id)
}

// GetState sets the state of a contract
func (k Keeper) GetType(ctx sdk.Context, id []byte) []byte {
	store := ctx.KVStore(k.typeStoreKey)
	return store.Get(id)
}

// GetContractsIterator returns an iterator over all stored contracts
func (k Keeper) GetContractsStateIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

func (k Keeper) GetContractsTypeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.typeStoreKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
