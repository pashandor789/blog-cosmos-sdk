package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	blog "github.com/blog"
)

var _ blog.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) blog.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetPost defines the handler for the Query/GetPost RPC method.
func (qs queryServer) GetPost(ctx context.Context, req *blog.QueryGetPostRequest) (*blog.QueryGetPostResponse, error) {
	post, err := qs.k.StoredPosts.Get(ctx, req.Index)
	if err == nil {
		return &blog.QueryGetPostResponse{Post: &post}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &blog.QueryGetPostResponse{Post: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
