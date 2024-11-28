package blog

import (
    "cosmossdk.io/errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (storedPost StoredPost) GetAddress() (black sdk.AccAddress, err error) {
    author, err := sdk.AccAddressFromBech32(storedPost.Author)
    return black, errors.Wrapf(err, ErrInvalidBlack.Error(), author)
}

func (storedPost StoredPost) Validate() (err error) {
    _, err = storedPost.GetAddress()
    if err != nil {
        return err
    }

    return err
}
