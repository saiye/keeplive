package main

import (
	"fmt"
	"game_go/report"
	"game_go/system"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := system.GetCfg("./")
	if err != nil {
		cfg, err = system.GetCfg("/etc/")
		if err != nil {
			fmt.Errorf("open keeplive config err:%v", err)
		}
		return
	}
	second := cfg.GetInt("app.second") // 定时器单位秒[1-N]
	if second == 0 {
		second = 60
	}
	// 创建一个定时器，每隔60秒触发一次
	ticker := time.NewTicker(time.Duration(second) * time.Second)
	// 在主函数退出前停止定时器
	defer ticker.Stop()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// 启动一个无限循环来处理定时事件
	for {
		select {
		case <-ticker.C:
			// 这里放置您想要定期执行的代码
			go func() {
				err2 := report.Report(cfg)
				if err2 != nil {
					fmt.Println(err2)
				}
			}()
		case <-quit:
			fmt.Println("The [keep live] service exits safely！")
			return
		}
	}
}
