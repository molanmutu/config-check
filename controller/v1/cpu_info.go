package v1

import (
	"config-check/settings"
	"config-check/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func rpmCPUCommand() (version string) {
	return "cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c"
}

func CpuInfo(c *gin.Context) {
	ip := *settings.Conf.LocalServerInfo
	for _, i2 := range ip.Servers {
		b := checkCPUInfo(i2)
		c.JSON(200, gin.H{
			i2: b,
		})
	}
}

// checkCPUInfo：获取cpu信息
func checkCPUInfo(ip string) (out string) {
	cli := utils.New(ip)
	fmt.Printf("%v\n", ip)
	command := rpmCPUCommand()
	output, err := cli.Run(command)
	if err != nil {
		zap.L().Error("获取cpu信息失败", zap.String("ip", ip), zap.String("command", command), zap.Error(err))
		return
	}
	mb := fmt.Sprintf("cpu info is :------%v", output)
	return mb

}
