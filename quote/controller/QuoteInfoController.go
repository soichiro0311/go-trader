package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
)

type QuoteInfoController struct {
	service *service.QuoteInfoService
}

func NewQuoteInfoController(service *service.QuoteInfoService) *QuoteInfoController {
	ctrl := new(QuoteInfoController)
	ctrl.service = service
	return ctrl
}

func (ctrl QuoteInfoController) InfosByCurPairHist(c *gin.Context) {
	cur1 := c.Param("cur1")
	cur2 := c.Param("cur2")
	c.JSON(200, ctrl.service.GetByCurrencyPair(cur1, cur2))
}

func (ctrl QuoteInfoController) InfosByCurPair(c *gin.Context) {
	cur1 := c.Param("cur1")
	cur2 := c.Param("cur2")
	c.JSON(200, ctrl.service.GetLatestByCurrencyPair(cur1, cur2))
}
