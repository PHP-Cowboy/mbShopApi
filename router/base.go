package router

import (
	"github.com/gin-gonic/gin"
	"mbShopApi/api"
)

func InitBaseRouter(g *gin.RouterGroup) {
	g.GET("captcha", api.GetCaptcha)
}
