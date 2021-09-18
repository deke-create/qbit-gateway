package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/deke-create/qbit-gateway/x/xfer/types"
)

var _ = strconv.Itoa(0)

func CmdAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action [id] [action]",
		Short: "Broadcast message action",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0])
			argsAction := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAction(clientCtx.GetFromAddress().String(), string(argsId), []byte(argsAction))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
