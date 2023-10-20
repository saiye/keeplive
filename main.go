package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	//cpu 和 内存超过该值报告消息给开发者
	var percent float64 = 1
	// 创建一个定时器，每隔1秒触发一次
	ticker := time.NewTicker(1 * time.Second)

	// 在主函数退出前停止定时器
	defer ticker.Stop()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// 启动一个无限循环来处理定时事件
	for {
		select {
		case <-ticker.C:
			// 这里放置您想要定期执行的代码
			ReportMessage(GetAllReportInfo(percent))
		case <-quit:
			fmt.Print("监控服务成功安全退出！")
			return
		}
	}
}
func GetAllReportInfo(minPercent float64) []string {
	infoList := make([]string, 0)
	m := ReportSystemMemory(minPercent)
	if m != "" {
		infoList = append(infoList, m)
	}
	cpu := ReportSystemCpu(minPercent)
	if cpu != "" {
		infoList = append(infoList, cpu)
	}
	return infoList
}
