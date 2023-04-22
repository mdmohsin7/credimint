package keeper

import (
	"context"
	"credimint/x/credimint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) LiquidStake(goCtx context.Context, msg *types.MsgLiquidStake) (*types.MsgLiquidStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// k.bankKeeper.Delegate(ctx, sdk.AccAddress(msg.Creator), 12, 0, staking.Validator{}, false)
	validator, err := k.StakingKeeper.GetValidator(ctx, sdk.ValAddress(msg.Validator))
	if err {
		panic(validator)
	}
	amount, err2 := sdk.ParseCoinNormalized(msg.Amount)
	if err2 != nil {
		panic(amount)
	}

	k.StakingKeeper.Delegate(ctx, sdk.AccAddress(msg.Creator), amount.Amount, staking.Unbonded, validator, true)
	return &types.MsgLiquidStakeResponse{}, nil
}
