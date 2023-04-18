package keeper

import (
	"context"
	"credimint/x/credimint/helper"
	"credimint/x/credimint/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amoount)
	var user, usrFound = k.GetUser(ctx, msg.Creator)
	if !usrFound {
		user = types.User{
			Index:               msg.Creator,
			CreditScore:         0,
			TimelyPayments:      0,
			DefaultRate:         strconv.Itoa(0),
			NumberOfLoans:       0,
			LoanDuration:        0,
			NumberOfLoansFunded: 0,
			LoanFundedDuration:  0,
			CollateralPercent:   strconv.Itoa(150),
		}
		k.SetUser(ctx, user)
	}
	user.NumberOfLoansFunded += 1
	helper.UpdateLenderScore(&user)
	k.SetUser(ctx, user)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	var borrowerUser, _ = k.GetUser(ctx, loan.Borrower)
	borrowerUser.NumberOfLoans += 1
	helper.UpdateBorrowerScore(&borrowerUser, &loan)
	k.SetUser(ctx, borrowerUser)
	loan.Lender = msg.Creator
	loan.State = "approved"
	loan.LenderScore = strconv.Itoa(int(user.CreditScore))
	k.SetLoan(ctx, loan)
	return &types.MsgApproveLoanResponse{}, nil
}
