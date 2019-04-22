package web

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Init(ctrl *controller.BoardInfoController) {
	r := router(ctrl)
	r.Run()
}

func router(ctrl *controller.BoardInfoController) *gin.Engine {
	r := gin.Default()

	u := r.Group("/board")
	{
		u.GET("/history/:cur1/:cur2", ctrl.InfosByCurPairHist)
		u.GET("/latest/:cur1/:cur2", ctrl.InfosByCurPair)
	}

	return r
}
