package keeper

import (
	"context"
	"math"
	"strconv"

	"credimint/x/credimint/helper"
	"credimint/x/credimint/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if msg.Creator != loan.Borrower {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the borrower")
	}
	amount, _ := sdk.ParseCoinsNormalized(loan.Amoount)
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	borrowerUser, _ := k.GetUser(ctx, loan.Borrower)
	lenderUser, _ := k.GetUser(ctx, loan.Lender)
	deadlineEpoch, _ := strconv.Atoi(loan.Deadline)
	repaymentTimeEpoch, _ := strconv.Atoi(msg.RepayTime)
	loanDuration := deadlineEpoch - repaymentTimeEpoch
	borrowerUser.LoanDuration += uint64(math.Abs(float64(loanDuration)))
	lenderUser.LoanFundedDuration += uint64(math.Abs(float64(loanDuration)))
	borrowerUser.TimelyPayments += 1
	helper.UpdateBorrowerScore(&borrowerUser, &loan)
	k.SetUser(ctx, borrowerUser)
	helper.UpdateLenderScore(&lenderUser)
	k.SetUser(ctx, lenderUser)
	loan.State = "repayed"
	k.SetLoan(ctx, loan)
	return &types.MsgRepayLoanResponse{}, nil
}
