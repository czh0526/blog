package keeper_test

import (
	"blog/x/blog/types"
	"testing"
)

func TestMsgServer_CreatePost(t *testing.T) {
	msgServer, goCtx := setupMsgServer(t)
	msgCreatePost := &types.MsgCreatePost{
		Creator: "Cai.Zhihong",
		Title:   "This is a test title.",
		Body:    "This is a test body.",
	}

	msgServer.CreatePost(goCtx, msgCreatePost)
}
