package keeper

import (
	"context"

	blog "github.com/blog"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *blog.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, indexedStoredPost := range data.IndexedStoredPostList {
		if err := k.StoredPosts.Set(ctx, indexedStoredPost.Index, indexedStoredPost.StoredPost); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*blog.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedPostList []blog.IndexedStoredPost
	if err := k.StoredPosts.Walk(ctx, nil, func(index string, storedPost blog.StoredPost) (bool, error) {
		indexedPostList = append(indexedPostList, blog.IndexedStoredPost{
			Index:      index,
			StoredPost: storedPost,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &blog.GenesisState{
		Params:                params,
		IndexedStoredPostList: indexedPostList,
	}, nil
}
