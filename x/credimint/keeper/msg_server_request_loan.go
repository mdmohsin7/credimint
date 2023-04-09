package keeper

import (
	"context"

	"credimint/x/credimint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"strconv"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var user, found = k.GetUser(ctx, msg.Creator)
	if !found {
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
	var riskLevel = "high"
	if user.CreditScore > 800 {
		riskLevel = "low"
	} else if user.CreditScore > 500 {
		riskLevel = "medium"
	} else if user.CreditScore < 500 {
		riskLevel = "high"
	}

	var loan = types.Loan{
		Amoount:       msg.Amount,
		Fee:           msg.Fee,
		Collateral:    msg.Collateral,
		Deadline:      msg.Deadline,
		State:         "requested",
		Borrower:      msg.Creator,
		RiskLevel:     riskLevel,
		BorrowerScore: strconv.Itoa(int(user.CreditScore)),
	}
	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}
	k.AppendLoan(ctx, loan)
	return &types.MsgRequestLoanResponse{}, nil
}
