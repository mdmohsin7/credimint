package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/credimint module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrUserNotFound         = sdkerrors.Register(ModuleName, 2, " User with this address not found")
	ErrWrongLoanState       = sdkerrors.Register(ModuleName, 3, "wrong loan state")
)
