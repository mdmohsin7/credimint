package keeper

import (
	"credimint/x/lqs/types"
)

var _ types.QueryServer = Keeper{}
