package blog

import "cosmossdk.io/errors"

var (
    ErrIndexTooLong     = errors.Register(ModuleName, 2, "index too long")
    ErrDuplicateAddress = errors.Register(ModuleName, 3, "duplicate address")
    ErrInvalidBlack     = errors.Register(ModuleName, 4, "author address is invalid: %s")
)
