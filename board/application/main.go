package main

import (
	"../controller"
	"../repository"
	"../service"
	"../web"
)

func main() {

	boardInfoRepository := repository.NewBoardInfoRepositoryImpl()
	boardInfoService := service.NewBoardInfoService(boardInfoRepository)
	boardInfoController := controller.NewBoardInfoController(boardInfoService)

	web.Init(boardInfoController)
}
