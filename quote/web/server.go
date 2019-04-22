package web

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Init(ctrl *controller.QuoteInfoController) {
	r := router(ctrl)
	r.Run()
}

func router(ctrl *controller.QuoteInfoController) *gin.Engine {
	r := gin.Default()

	u := r.Group("/quote")
	{
		u.GET("/history/:cur1/:cur2", ctrl.InfosByCurPairHist)
		u.GET("/latest/:cur1/:cur2", ctrl.InfosByCurPair)
	}

	return r
}
