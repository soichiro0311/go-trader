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
	boardInfoRepository := repository.NewBoardInfoRepositoryImpl()
	orderService := service.NewOrderService(orderRepository, boardInfoRepository)
	orderController := controller.NewOrderController(orderService)

	web.Init(orderController)
}
