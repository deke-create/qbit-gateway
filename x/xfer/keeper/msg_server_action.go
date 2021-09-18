package keeper

import (
	"context"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

func (k msgServer) Action(goCtx context.Context, msg *types.MsgAction) (*types.MsgActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//TODO:  Need to handle results and events more thoroughly
	_, cErr := k.handleMsgExecuteAction(ctx, *msg)
	return &types.MsgActionResponse{
		Id: "-0",
	}, cErr
}


func (k msgServer) handleMsgExecuteAction(ctx sdk.Context, msg types.MsgAction) (*sdk.Result, error) {
	connection := k.GetConnection()
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	bh := make([]byte, 8)
	binary.LittleEndian.PutUint64(bh, uint64(k.BaseApp.LastBlockHeight()))
	actionResult, err := connection.SendAction(ctx, addr, msg.Action, bh)
	if err != nil {
		//return sdk.ErrInternal("Couldn't execute action!").Result()
		return nil, err
	}
	return &sdk.Result{Events: actionResult.Events}, nil

}