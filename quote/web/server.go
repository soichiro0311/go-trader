package web

import (
	"../controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(ctrl *controller.QuoteInfoController) {
	r := router(ctrl)
	r.Run(":8091")
}

func router(ctrl *controller.QuoteInfoController) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	u := r.Group("/quote")
	{
		u.GET("/latest/:cur1/:cur2", ctrl.InfosByCurPair)
	}

	return r
}
