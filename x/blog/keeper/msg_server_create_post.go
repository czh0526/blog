package keeper

import (
	"context"
	"fmt"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	post := types.Post{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Title:     msg.Title,
		Body:      msg.Body,
		CreatedAt: ctx.BlockHeight(),
	}

	fmt.Println("k.AppendPost() was called.")
	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
