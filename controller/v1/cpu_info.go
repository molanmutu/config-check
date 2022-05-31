package v1

import (
	"config-check-g4/settings"
	"config-check-g4/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//获取命令
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

func checkCPUInfo(ip string) (out string) {
	cli := utils.New(ip)
	fmt.Printf("%v\n", ip)
	command := rpmCPUCommand()
	output, err := cli.Run(command)
	if err != nil {
		fmt.Println("Get cpu info failed: ", err)
	}
	mb := fmt.Sprintf("cpu info is :------%v", output)
	return mb

}
