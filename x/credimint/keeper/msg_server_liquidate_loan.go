package keeper

import (
	"context"
	"strconv"

	"credimint/x/credimint/helper"
	"credimint/x/credimint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Lender != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
	}
	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	deadlineEpoch, err := strconv.Atoi(loan.Deadline)
	if err != nil {
		panic(err)
	}
	liquidationTimeEpoch, err := strconv.Atoi(msg.LiquidationTime)
	if err != nil {
		panic(err)
	}
	if liquidationTimeEpoch < deadlineEpoch {
		return nil, sdkerrors.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "liquidated"
	borrowerUser, found := k.GetUser(ctx, loan.Borrower)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %s doesn't exist", loan.Borrower)
	}
	lenderUser, found := k.GetUser(ctx, loan.Lender)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %s doesn't exist", loan.Lender)
	}
	loanDuration := liquidationTimeEpoch - deadlineEpoch
	borrowerUser.LoanDuration += uint64(loanDuration)
	lenderUser.LoanFundedDuration += uint64(loanDuration)
	helper.UpdateBorrowerScore(&borrowerUser, &loan)
	k.SetUser(ctx, borrowerUser)
	helper.UpdateLenderScore(&lenderUser)
	k.SetUser(ctx, lenderUser)
	k.SetLoan(ctx, loan)
	return &types.MsgLiquidateLoanResponse{}, nil
}
