package helper

import (
	"credimint/x/credimint/types"
)

func UpdateCollateralPercentage(user *types.User) {
	// Initialize the credit score to the current credit score of the lender
	CollateralPercent := user.CollateralPercent

	// Adjust the collateral percentage based on the credit score
	if user.CreditScore > 800 {
		CollateralPercent = "50"
	}
	if user.CreditScore > 700 {
		CollateralPercent = "60"
	}
	if user.CreditScore > 600 {
		CollateralPercent = "70"
	}
	if user.CreditScore > 500 {
		CollateralPercent = "80"
	}
	if user.CreditScore > 400 {
		CollateralPercent = "90"
	}
	if user.CreditScore > 300 {
		CollateralPercent = "100"
	}
	if user.CreditScore > 200 {
		CollateralPercent = "150"
	}
	if user.CreditScore > 100 {
		CollateralPercent = "200"
	}

	// Set the updated collateral percentage for the user
	user.CollateralPercent = CollateralPercent

}
