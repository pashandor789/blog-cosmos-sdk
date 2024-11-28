package blog

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

    unique := make(map[string]bool)
    for _, indexedStoredPost := range gs.IndexedStoredPostList {
        if length := len([]byte(indexedStoredPost.Index)); MaxIndexLength < length || length < 1 {
            return ErrIndexTooLong
        }
        if _, ok := unique[indexedStoredPost.Index]; ok {
            return ErrDuplicateAddress
        }
        if err := indexedStoredPost.StoredPost.Validate(); err != nil {
            return err
        }
        unique[indexedStoredPost.Index] = true
    }

	return nil
}
