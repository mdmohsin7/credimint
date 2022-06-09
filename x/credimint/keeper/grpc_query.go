package keeper

import (
	"credimint/x/credimint/types"
)

var _ types.QueryServer = Keeper{}
