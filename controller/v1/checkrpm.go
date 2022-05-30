package v1

import (
	"config-check/settings"
	"config-check/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取命令
func rpmCommand() (version string) {
	return "rpm -q nmap-ncat"
}

// CheckRPM 返回json给前端
func CheckRPM(c *gin.Context) {
	ip := *settings.Conf.LocalServerInfo
	for _, i2 := range ip.Servers {
		b := checkRpmInfo(i2)
		c.JSON(200, gin.H{
			i2: b,
		})
	}
}

//拷贝rpm包并安装
func checkRpmInfo(ip string) (out string) {
	cli := utils.New(ip)
	fmt.Printf("%v\n", ip)
	command := rpmCommand()
	output, err := cli.Run(command)
	if err != nil {
		ScpRPM(cli)
		out, err := cli.Run("yum install -y " + cli.Des + "/*.rpm")
		if err != nil {
			fmt.Println("yum install rpm err:", err)
			zap.L().Error("yum install rpm err", zap.Error(err))
			mm := fmt.Sprintf("RPM %v  installed err:%v", command, err)
			return mm
		}
		fmt.Println("yum install rpm info:", out)
		output2, _ := cli.Run(command)
		mb := fmt.Sprintf("RPM sshpass is installing,veriosn is :------%v", output2)
		return mb
	}
	mb := fmt.Sprintf("RPM sshpass is  installed,veriosn is :------%v", output)
	return mb
}
