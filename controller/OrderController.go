package controller

import (
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *service.OrderService
}

func NewOrderController(service *service.OrderService) *OrderController {
	ctrl := new(OrderController)
	ctrl.service = service
	return ctrl
}

func (ctrl OrderController) RegisterOrder(c *gin.Context) {
	req := model.RegisterOrderRequest{}
	c.BindJSON(&req)
	ctrl.service.RegisterOrder(req)
}
