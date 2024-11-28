package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	blog "github.com/blog"
)

type msgServer struct {
	k Keeper
}

var _ blog.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) blog.MsgServer {
	return &msgServer{k: keeper}
}

// CreatePost defines the handler for the MsgCreatePost message.
func (ms msgServer) CreatePost(ctx context.Context, msg *blog.MsgCreatePost) (*blog.MsgCreatePostResponse, error) {
	if length := len([]byte(msg.Index)); blog.MaxIndexLength < length || length < 1 {
		return nil, blog.ErrIndexTooLong
	}
	if _, err := ms.k.StoredPosts.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("post already exists at index: %s", msg.Index)
	}

	storedpost := blog.StoredPost{
		Text:   msg.Text,
		Author: msg.Creator,
	}
	if err := storedpost.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredPosts.Set(ctx, msg.Index, storedpost); err != nil {
		return nil, err
	}

	return &blog.MsgCreatePostResponse{}, nil
}
