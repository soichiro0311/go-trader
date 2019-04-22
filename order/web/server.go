package web

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Init(ctrl *controller.OrderController) {
	r := router(ctrl)
	r.Run(":8090")
}

func router(ctrl *controller.OrderController) *gin.Engine {
	r := gin.Default()

	u := r.Group("/order")
	{
		u.POST("/register", ctrl.RegisterOrder)
		u.GET("/all", ctrl.Orders)
	}

	return r
}
