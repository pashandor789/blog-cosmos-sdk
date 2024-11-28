package blog

import "cosmossdk.io/collections"

const ModuleName = "blog"
const MaxIndexLength = 256

var (
	ParamsKey      = collections.NewPrefix("Params")
	StoredPostsKey = collections.NewPrefix("StoredPost/value/")
)
