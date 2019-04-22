package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
)

type BoardInfoController struct {
	service *service.BoardInfoService
}

func NewBoardInfoController(service *service.BoardInfoService) *BoardInfoController {
	ctrl := new(BoardInfoController)
	ctrl.service = service
	return ctrl
}

func (ctrl BoardInfoController) InfosByCurPairHist(c *gin.Context) {
	cur1 := c.Param("cur1")
	cur2 := c.Param("cur2")
	c.JSON(200, ctrl.service.GetByCurrencyPair(cur1, cur2))
}

func (ctrl BoardInfoController) InfosByCurPair(c *gin.Context) {
	cur1 := c.Param("cur1")
	cur2 := c.Param("cur2")
	c.JSON(200, ctrl.service.GetLatestByCurrencyPair(cur1, cur2))
}
