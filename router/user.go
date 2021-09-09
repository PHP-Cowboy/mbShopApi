package router

import (
	"github.com/gin-gonic/gin"
	"mbShopApi/api"
)

func InitUserRouter(r *gin.RouterGroup) {
	r.GET("list", api.GetUserList)
}
