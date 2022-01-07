package keeper

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

type msgServer struct {
	Keeper
	BaseApp           baseapp.BaseApp
	XferServerAddress string
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, bApp baseapp.BaseApp, xferServerAddress string) types.MsgServer {

	return &msgServer{Keeper: keeper, XferServerAddress: xferServerAddress, BaseApp: bApp}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) GetConnection() Connection {
	return NewConnection(k.XferServerAddress, k.Keeper)
}
