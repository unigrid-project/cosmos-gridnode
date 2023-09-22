package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	"github.com/spf13/cobra"
	gridnode "github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types/gridnode"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/version"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
	FlagMetadata               = "metadata"
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

	cmd.AddCommand(GetCmdDelegate())
	cmd.AddCommand(GetCmdQueryDelegatedAmount())

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

func NewCmdCastVoteFromGridnode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cast-vote [proposal-id] [option]",
		Short: "Cast a vote from a Gridnode, options: yes/no/no_with_veto/abstain",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a vote for an active proposal. You can
find the proposal-id by running "%s query gov proposals".

Example:
$ %s tx gov vote 1 yes --from mykey
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// Get voting address
			from := clientCtx.GetFromAddress()

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			metadata, err := cmd.Flags().GetString(FlagMetadata)
			if err != nil {
				return err
			}

			// Find out which vote option user chose
			byteVoteOption, err := v1.VoteOptionFromString(govutils.NormalizeVoteOption(args[1]))
			if err != nil {
				return err
			}

			// Build vote message and run basic validation
			msg := v1.NewMsgVote(from, proposalID, byteVoteOption, metadata)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagMetadata, "", "Specify metadata of the vote")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryDelegatedAmount() *cobra.Command {
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
			client := gridnode.NewGridnodeQueryClient(conn)

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
