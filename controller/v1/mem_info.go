package v1

import (
	"config-check/settings"
	"config-check/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//获取命令
func memCommand() (version string) {
	return "free -h|grep 'Mem:'|awk '{print $2}'"
}

// MemInfo 获取内存信息
func MemInfo(c *gin.Context) {
	ip := *settings.Conf.LocalServerInfo
	for _, i2 := range ip.Servers {
		b := checkMenInfo(i2)
		c.JSON(200, gin.H{
			i2: b,
		})
	}
}

// checkMenInfo 检查内存信息
func checkMenInfo(ip string) (out string) {
	cli := utils.New(ip)
	fmt.Printf("%v\n", ip)
	command := memCommand()
	output, err := cli.Run(command)
	if err != nil {
		zap.L().Error("获取内存信息失败", zap.String("ip", ip), zap.String("command", command), zap.Error(err))
		return
	}
	mb := fmt.Sprintf("men info is :------%v", output)
	return mb
}
