package cli

import (
	"context"
	"github.com/spf13/cast"
	"strconv"

	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-comment",
		Short: "list all comment",
		RunE: func(cmd *cobra.Command, args []string) error {
			argPostId, err := cast.ToUint64E(args[0])
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCommentsRequest{
				PostID:     argPostId,
				Pagination: pageReq,
			}

			res, err := queryClient.Comments(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-comment [id]",
		Short: "shows a comment",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetCommentRequest{
				Id: id,
			}

			res, err := queryClient.Comment(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
