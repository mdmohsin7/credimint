package keeper

import (
	"context"

	"credimint/x/lqs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) Liquid(goCtx context.Context, msg *types.MsgLiquid) (*types.MsgLiquidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	delegater, _ := sdk.AccAddressFromBech32(msg.Creator)
	validator, _ := sdk.ValAddressFromBech32(msg.Validator)
	val, _ := k.stakingKeeper.GetValidator(ctx, validator)
	amount, _ := sdk.ParseCoinNormalized(msg.Amount)
	sTokens := sdk.NewCoin("scred", amount.Amount)
	k.stakingKeeper.Delegate(ctx, delegater, amount.Amount, stakingtypes.Unbonded, val, true)
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sTokens))
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, delegater, sdk.NewCoins(sTokens))
	_ = ctx

	return &types.MsgLiquidResponse{}, nil
}
