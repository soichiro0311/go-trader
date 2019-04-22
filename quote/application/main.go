package main

import (
	"../controller"
	"../repository"
	"../service"
	"../web"
)

func main() {

	quoteInfoRepository := repository.NewQuoteInfoRepositoryImpl()
	quoteInfoService := service.NewQuoteInfoService(quoteInfoRepository)
	quoteInfoController := controller.NewQuoteInfoController(quoteInfoService)

	web.Init(quoteInfoController)
}
