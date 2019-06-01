package controller

import (
	"../dto"
	"../enum"
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

func (ctrl QuoteInfoController) InfosByCurPair(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Max-Age", "86400")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	cur1 := c.Param("cur1")
	cur2 := c.Param("cur2")
	latestInfo := ctrl.service.GetLatestByCurrencyPair(cur1, cur2)
	latestInfoDto := map[enum.SideEnum]dto.QuoteInfoDto{}
	for i, v := range latestInfo {
		latestInfoDto[i] = *dto.NewQuoteInfoDto(v)
	}
	c.JSON(200, latestInfoDto)
}
