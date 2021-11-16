package main

import (
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/modules/transactions/transport/pbtransaction"
	"nolan/spin-game/modules/wallets/transport/pbwallet"
)

func setupPubsubRoute(appCtx appctx.AppContext) {

	{
		pbwallet.CreateWallet(appCtx)
	}

	{
		pbtransaction.RequestDeposit(appCtx)
	}

}