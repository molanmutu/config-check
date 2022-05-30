package routers

import (
	v1 "config-check/controller/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//v1版本
	GroupV1 := r.Group("/api/v1")
	{
		GroupV1.GET("/ping", v1.Ping)
		GroupV1.GET("/rpminfo", v1.CheckRPM)
	}

	//v1.GET("/Allinfo",controller.AllInfo)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})

	})
	return r
}
