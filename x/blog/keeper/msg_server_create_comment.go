package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateComment(goCtx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d does not exist", msg.PostID))
	}

	var comment = types.Comment{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Body:      msg.Body,
		Title:     msg.Title,
		PostID:    post.Id,
		CreatedAt: ctx.BlockHeight(),
	}

	//if comment.CreatedAt > post.CreatedAt+100 {
	//	return nil, sdkerrors.Wrapf(types.ErrCommentOld, "Comment created at %d is older than post created at %d",
	//		comment.CreatedAt, post.CreatedAt)
	//}

	id := k.AppendComment(ctx, comment)

	return &types.MsgCreateCommentResponse{Id: id}, nil
}
