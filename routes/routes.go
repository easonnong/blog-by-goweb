package routes

import (
	"net/http"

	"github.com/easonnong/blog-by-goweb/utils"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	//引入配置
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//不同版本v1 v2
	router := r.Group("api/v1")
	{
		router.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
