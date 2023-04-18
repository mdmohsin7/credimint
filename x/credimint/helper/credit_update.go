package helper

import (
	"credimint/x/credimint/types"
	"strconv"
)

func UpdateLenderScore(lender *types.User) {
	// Initialize the credit score to the current credit score of the lender
	creditScore := lender.CreditScore

	// Adjust the credit score based on the number of loans funded
	creditScore += lender.NumberOfLoansFunded * 10

	// Adjust the credit score based on the duration of loans funded
	creditScore += lender.LoanFundedDuration / 5

	// Check if the credit score is within the range of 0 to 1000
	if creditScore < 0 {
		creditScore = 0
	} else if creditScore > 1000 {
		creditScore = 1000
	}

	// Set the updated credit score for the lender
	lender.CreditScore = creditScore
}

func UpdateBorrowerScore(borrower *types.User, loan *types.Loan) {
	// Initialize the credit score to the current credit score of the borrower
	creditScore := borrower.CreditScore

	// Adjust the credit score based on the age of the accounts
	// creditScore += borrower.AccountAge / 5

	// Adjust the credit score based on the number of timely payments
	creditScore += borrower.TimelyPayments * 10

	var defaultedLoansCount = borrower.NumberOfLoans - borrower.TimelyPayments
	var defaultRate = defaultedLoansCount / borrower.NumberOfLoans
	// Adjust the credit score based on the default rate
	creditScore -= uint64(defaultRate) * 10

	// Adjust the credit score based on the number of loans
	creditScore += borrower.NumberOfLoans * 10

	// Adjust the credit score based on the duration of loans
	creditScore += borrower.LoanDuration / 5

	// Adjust the credit score based on the collateral
	var loanCollateral, _ = strconv.ParseFloat(loan.Collateral, 64)
	creditScore += uint64(loanCollateral) / 5

	// Adjust the credit score based on the fee
	// var loanFee, _ = strconv.ParseFloat(loan.Fee, 64)
	// creditScore -= uint64(loanFee) * 10

	// Check if the credit score is within the range of 0 to 1000
	if creditScore < 0 {
		creditScore = 0
	} else if creditScore > 1000 {
		creditScore = 1000
	}

	// Set the updated credit score for the borrower
	borrower.CreditScore = creditScore
}
