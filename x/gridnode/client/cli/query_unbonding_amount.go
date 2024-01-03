package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func CmdUnbondingEntries() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbonding-entries [delegator-address]",
		Short: "Query unbonding entries for an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			bondingAccountAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryUnbondingEntriesRequest{
				BondingAccountAddress: bondingAccountAddress,
			}

			res, err := queryClient.UnbondingEntries(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
