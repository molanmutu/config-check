package main

import (
	"config-check/logger"
	"config-check/routers"
	"config-check/settings"
	"fmt"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed，err:%v\n", err)
		return
	}

	r := routers.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
		return
	}
}
