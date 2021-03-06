package v1

import (
	"config-check/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os/exec"
)

func Ping(c *gin.Context) {
	ip := *settings.Conf.LocalServerInfo
	for _, i2 := range ip.Servers {
		b := netWorkStatus(i2)
		c.JSON(200, gin.H{
			i2: b,
		})
	}
}

// netWorkStatus 检测服务器状态是否正常
func netWorkStatus(ip string) bool {
	cmd := exec.Command("ping", ip)
	err := cmd.Run()
	if err != nil {
		zap.L().Error("Net comm filed", zap.Error(err))
		return false
	} else {
		fmt.Println("Net Status , OK")
	}
	return true
}
