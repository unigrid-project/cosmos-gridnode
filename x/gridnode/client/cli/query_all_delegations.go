package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/unigrid-project/cosmos-sdk-gridnode/x/gridnode/types"
)

func CmdAllDelegations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-delegations",
		Short: "Query all delegations",
		Args:  cobra.NoArgs, // No arguments required for this command
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDelegationsRequest{
				// Optionally, you can add pagination parameters here if needed
				// Pagination: &types.Pagination{Limit: 10, Offset: 0},
			}

			res, err := queryClient.AllDelegations(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
