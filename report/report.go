package report

import (
	"errors"
	"fmt"
	"game_go/request"
	"game_go/system"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

// Report 报告系统状态
func Report(cfg *viper.Viper) error {
	//报告系统消息
	err := ReportSystemInfo(cfg)
	if err != nil {
		return err
	}
	//报告服务状态
	return nil
}

// ReportHttpServiceInfo 报告服务状态
func ReportHttpServiceInfo(cfg *viper.Viper) error {

	count := cfg.GetInt("check.count") // url 数量为n
	res := make([]string, 0)
	if count > 0 {
		for i := 1; i <= count; i++ {
			url := cfg.GetString(fmt.Sprintf("check.url%d", i))
			if url != "" {
				res = append(res, url)
			}
		}
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	var wg = &sync.WaitGroup{}

	output := make([]string, 0)
	ch := make(chan string, 2)
	for _, url := range res {
		wg.Add(1)
		go func(url string) {
			resp, statusCode, err := request.HTTPRequest(request.POST, url, strings.NewReader("{}"), &header)
			var s string
			if statusCode != 200 || err != nil {
				if err != nil {
					s = err.Error()
				}
				ch <- fmt.Sprintf("statusCode:%d,err:%s,resp:%s", statusCode, s, resp)
			}
			defer wg.Done()
		}(url)
	}
	output = append(output, <-ch)
	defer func() {
		wg.Wait()
		close(ch)
	}()
	if len(output) > 0 {
		return errors.New(strings.Join(output, "\n"))
	}
	return nil
}

// ReportSystemInfo 报告系统消息
func ReportSystemInfo(cfg *viper.Viper) error {
	env := cfg.GetString("app.env")          // 读取配置
	percent := cfg.GetFloat64("app.percent") // 警告百分比0-100
	messageList := GetAllReportInfo(percent)
	phoneList := GetPhoneList(cfg.GetString("dingtalk.phone_list"))
	keyword := cfg.GetString("dingtalk.keyword") // 读取警报关键词
	message := ""
	if percent > 95 {
		percent = 95
	}
	for _, msg := range messageList {
		receiver := &TextMessage{
			At: AtContent{
				AtMobiles: phoneList,
				AtUserIds: nil,
				IsAtAll:   len(phoneList) == 0,
			},
			Text: TextContent{
				Content: fmt.Sprintf("【环境:%s】%s:%s", env, keyword, msg),
			},
			MsgType: "text",
		}
		res := receiver.Send(cfg)
		if res != nil {
			message += res.Error()
		}
	}
	if message != "" {
		return errors.New(message)
	}
	return nil
}

func GetPhoneList(str string) []string {
	return strings.Split(str, ",")
}
func GetAllReportInfo(minPercent float64) []string {
	infoList := make([]string, 0)
	m := system.ReportSystemMemory(minPercent)
	if m != "" {
		infoList = append(infoList, m)
	}
	cpu := system.ReportSystemCpu(minPercent)
	if cpu != "" {
		infoList = append(infoList, cpu)
	}
	return infoList
}
