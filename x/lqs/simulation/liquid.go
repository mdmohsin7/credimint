package simulation

import (
	"math/rand"

	"credimint/x/lqs/keeper"
	"credimint/x/lqs/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgLiquid(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLiquid{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Liquid simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Liquid simulation not implemented"), nil, nil
	}
}
