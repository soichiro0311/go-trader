package main

import (
	"../controller"
	"../db"
	"../repository"
	"../service"
	"../web"
)

func main() {
	db.Init()

	orderRepository := repository.NewOrderRepositoryImpl()
	quoteInfoRepository := repository.NewQuoteInfoRepositoryImpl()
	orderService := service.NewOrderService(orderRepository, quoteInfoRepository)
	orderController := controller.NewOrderController(orderService)

	web.Init(orderController)
}
