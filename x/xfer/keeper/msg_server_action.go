package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

func (k msgServer) Action(goCtx context.Context, msg *types.MsgAction) (*types.MsgActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgActionResponse{}, nil
}
