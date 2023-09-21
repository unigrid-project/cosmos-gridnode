package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	gridnode "github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types/gridnode"
	"google.golang.org/grpc"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	return cmd
}

func GetCmdDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate",
		Short: "Delegate tokens for gridnode",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := cmd.Flags().GetInt64("amount")
			if err != nil {
				return err
			}

			// Create the message
			delegatorAddress := clientCtx.GetFromAddress()
			msg := types.NewMsgDelegateGridnode(delegatorAddress, amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Int64("amount", 0, "Amount of tokens to delegate")
	cmd.MarkFlagRequired("amount")

	return cmd
}

func GetCmdQueryDelegatedAmount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delegated-amount [delegator-address]",
		Short: "Query the amount delegated by the specified account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			delegatorAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// Set up the gRPC client
			conn, err := grpc.Dial(clientCtx.NodeURI, grpc.WithInsecure())
			if err != nil {
				return err
			}
			defer conn.Close()
			client := gridnode.NewQueryClient(conn)

			// Make the gRPC request
			req := &gridnode.QueryDelegatedAmountRequest{
				DelegatorAddress: delegatorAddr.String(),
			}
			res, err := client.DelegatedAmount(context.Background(), req)
			if err != nil {
				return err
			}
			// Handle the response
			amount := res.Amount
			response := &gridnode.QueryDelegatedAmountResponse{
				Amount: amount,
			}

			return clientCtx.PrintProto(response)
		},
	}
}
