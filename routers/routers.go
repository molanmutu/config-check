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
		GroupV1.GET("/ping", v1.Ping)        //测试服务器是否通信
		GroupV1.GET("/rpminfo", v1.CheckRPM) //检查rpm包是否存在
		GroupV1.GET("/cpuinfo", v1.CpuInfo)  //cpu信息
		GroupV1.GET("/meminfo", v1.MemInfo)  //内存信息
	}

	//v1.GET("/Allinfo",controller.AllInfo)

	//
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})

	})
	return r
}
