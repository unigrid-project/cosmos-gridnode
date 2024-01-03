package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

var _ = strconv.Itoa(0)

func CmdDelegatedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegated-amount [delegator-address]",
		Short: "Query delegated-amount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDelegatorAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDelegatedAmountRequest{

				DelegatorAddress: reqDelegatorAddress,
			}

			res, err := queryClient.DelegatedAmount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
