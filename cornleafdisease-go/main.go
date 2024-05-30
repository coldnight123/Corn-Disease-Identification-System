package main

import (
	"fmt"
	"project/cornleafdisease/controller"
	"project/cornleafdisease/logger"
	"project/cornleafdisease/pkg/snowflake"
	"project/cornleafdisease/router"
	"project/cornleafdisease/setting"
)

func main() {
	// 加载配置
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err: %v\n", err)
		return
	}
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
