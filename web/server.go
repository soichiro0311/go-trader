package web

import (
	"github.com/gin-gonic/gin"
)

func Init(ctrl *controller.OrderController) {
	r := router(ctrl)
	r.Run()
}

func router(ctrl *controller.OrderController) *gin.Engine {
	r := gin.Default()

	u := r.Group("/order")
	{
		u.POST("/register", ctrl.RegisterOrder)
	}

	return r
}
