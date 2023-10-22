package main

import (
	"fmt"
	"game_go/report"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 创建一个定时器，每隔60秒触发一次
	ticker := time.NewTicker(60 * time.Second)
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
				err := report.ReportMessage("./")
				if err != nil {
					fmt.Println(err)
				}
			}()
		case <-quit:
			fmt.Println("The [keep live] service exits safely！")
			return
		}
	}
}
