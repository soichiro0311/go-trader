package main

import (
	"../component"
	"../controller"
	"../db"
	"../repository"
	"../service"
	"../web"
)

func main() {
	db.Init()

	orderRepository := repository.NewOrderRepositoryImpl()
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)
	component.GetPriceGenerator().Generate("BTC", "USD")

	web.Init(orderController)
}
