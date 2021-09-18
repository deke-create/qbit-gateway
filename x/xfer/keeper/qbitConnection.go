package keeper

import (
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deke-create/qbit-gateway/proto/proto_vendor"
	types2 "github.com/tendermint/tendermint/abci/types"

	"io"

	"google.golang.org/grpc"
)

// Connection represents a connection to the .NET StrongForce app
type Connection struct {
	serverAddress string
	keeper        Keeper

	client  proto_vendor.AppConnectClient
	channel *grpc.ClientConn
}

// NewConnection creates a new connection to the .NET  app
func NewConnection(serverAddress string, keeper Keeper) Connection {
	return Connection{
		serverAddress: serverAddress,
		keeper:        keeper,
	}
}

func (c *Connection) ensureConnected() bool {
	if c.client == nil {
		var err error
		c.channel, err = grpc.Dial(c.serverAddress, grpc.WithInsecure())
		if err != nil {
			println("Error while connecting: " + err.Error())
			return false
		}
		c.client = proto_vendor.NewAppConnectClient(c.channel)
	}
	return true
}

// SendAction sends an action over the connection
func (c *Connection) SendAction(ctx types.Context, from types.AccAddress, action []byte, blockHeight []byte) (*types.Result, error) {
	if !c.ensureConnected() {
		panic("Could not connect to the .NET app!")
	}
	goContext := ctx.Context()
	stream, err := c.client.ExecuteAction(goContext)
	if err != nil {
		println("Error while communicating to .NET: " + err.Error())
		//return types.Result{Code: types.CodeInternal, Codespace: "Error while communicating to .NET"}
		//errMsg := fmt.Sprintf("failed node validation %s message type: %T", types.ModuleName, msg)
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Error while communicating to .NET")
	}

	waitc := make(chan *types.Result)
	events := types.EmptyEvents()
	go func() {
		defer func() {
			if r := recover(); r != nil {

				events.AppendEvent(
					types.Event{
						Type: "ErrorConnectingToDotNet",
					},
				)
				waitc <- &types.Result{Events: events.ToABCIEvents(), Log: "Error while communicating to .NET"}
				//types.Result{Code: types.CodeInternal, Codespace: "Error while communicating to .NET"}
				close(waitc)
				stream.CloseSend()
				return
			}
		}()

		for {
			request, err := stream.Recv()
			if err == io.EOF {
				waitc <- &types.Result{Events: events.ToABCIEvents()}
				close(waitc)
				stream.CloseSend()
				return
			}
			if err != nil {
				println("Error while communicating to .NET: " + err.Error())
				waitc <- &types.Result{Events: events.ToABCIEvents(), Log: "Error while communicating to .NET"}
				close(waitc)
				return
			}
			if request.Data == nil {
				err = stream.Send(&proto_vendor.ActionOrContract{
					Data: &proto_vendor.ActionOrContract_Contract{
						&proto_vendor.ContractResponse{
							Address: request.Address,
							Data:    c.keeper.GetState(ctx, request.Address),
						},
					},
				})
				if err != nil {
					println("Error while communicating to .NET: " + err.Error())
					waitc <- &types.Result{Events: events.ToABCIEvents(), Log: "Error while communicating to .NET"}
					close(waitc)
					return
				}
			} else {
				events = events.AppendEvent(
					types.Event{
						Type: "ContractStateUpdate",
						Attributes: []types2.EventAttribute{
							{Key: []byte("Address"), Value: []byte(base64.RawURLEncoding.EncodeToString(request.Address))},
							{Key: []byte("State"), Value: request.Data},
						},
					},
				)
				c.keeper.SetState(ctx, request.Address, request.Data, request.TypeName)
			}
		}
	}()

	if err := stream.Send(&proto_vendor.ActionOrContract{
		Data: &proto_vendor.ActionOrContract_Action{
			&proto_vendor.XAction{
				Address:     from,
				Data:        action,
				Blockheight: blockHeight,
			},
		},
	}); err != nil {
		println("Error while communicating to .NET: " + err.Error())
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Error while communicating to .NET")
	}
	return <-waitc, nil
}
