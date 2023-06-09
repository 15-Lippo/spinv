package gingamespin

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	gamespinbiz "nolan/spin-game/modules/game_spin/biz"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
	gamespinstorage "nolan/spin-game/modules/game_spin/storage"
	walletstorage "nolan/spin-game/modules/wallets/storage"

	"github.com/gin-gonic/gin"
)

func BuyPackage(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data gamespinmodel.BuyPackage
		user := c.MustGet(common.CurrentUser).(*tokenprovider.TokenPayload)
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.Fullfill()

		db := appctx.GetMaiDBConnection()
		walletStorage := walletstorage.NewSQLStore(db)
		gameSpinStorage := gamespinstorage.NewSQLStore(db)
		gameSpinBiz := gamespinbiz.NewBuyPackageBiz(walletStorage, gameSpinStorage)

		result, err := gameSpinBiz.BuyPackage(c, user, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
