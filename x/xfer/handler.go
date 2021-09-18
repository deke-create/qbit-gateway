package xfer

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/baseapp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deke-create/qbit-gateway/x/xfer/keeper"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper, serverAddress string, baseapp baseapp.BaseApp) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k, baseapp, serverAddress)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgAction:
			res, err := msgServer.Action(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
