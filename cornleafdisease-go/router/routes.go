package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/cornleafdisease/controller"
	"project/cornleafdisease/logger"
	"project/cornleafdisease/middlewares"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middlewares.Cors,
	)
	r.POST("/signup", controller.SignupHandler)
	r.POST("/login", controller.LoginHandler)
	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.POST("/record", controller.GetRecordListHandler)
		r.GET("/search", controller.GetSearchHandler)
		r.POST("/upImg", controller.ImgAnalysisHandler)
		r.GET("/count", controller.GetCountHandler)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
