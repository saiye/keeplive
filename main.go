package main

import "game_go/report"

func main() {
	report.ReportMessage("./")
	//// 创建一个定时器，每隔2秒触发一次
	//ticker := time.NewTicker(2 * time.Second)
	//// 在主函数退出前停止定时器
	//defer ticker.Stop()
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//// 启动一个无限循环来处理定时事件
	//for {
	//	select {
	//	case <-ticker.C:
	//		// 这里放置您想要定期执行的代码
	//		report.ReportMessage("./")
	//	case <-quit:
	//		fmt.Print("The [keep live] service exits safely！")
	//		return
	//	}
	//}
}
