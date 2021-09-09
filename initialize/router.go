package initialize

import (
	"github.com/gin-gonic/gin"
	"mbShopApi/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/u/v1")

	router.InitUserRouter(apiGroup)

	return r
}
