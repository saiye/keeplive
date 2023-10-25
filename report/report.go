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

type UrlInfo struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

// Report 报告系统状态
func Report(cfg *viper.Viper) error {
	var msg string
	//报告系统消息
	err := SystemInfo(cfg)
	if err != nil {
		msg += err.Error()
	}
	//报告服务状态
	err2 := HttpServiceInfo(cfg)
	if err2 != nil {
		msg += err2.Error()
	}
	if msg != "" {
		return errors.New(msg)
	}
	return nil
}

// HttpServiceInfo 报告服务状态
func HttpServiceInfo(cfg *viper.Viper) error {
	messageList := GetHttpServiceErrorInfo(cfg)
	if len(messageList) > 0 {
		//do report
		err := SendTextMessage(cfg, messageList)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHttpServiceErrorInfo 报告服务状态
func GetHttpServiceErrorInfo(cfg *viper.Viper) []string {
	count := cfg.GetInt("check.count") // url 数量为n
	res := make([]*UrlInfo, 0)
	if count > 0 {
		for i := 1; i <= count; i++ {
			url := cfg.GetString(fmt.Sprintf("check.url%d", i))
			if url != "" {
				res = append(res, &UrlInfo{
					Url:  url,
					Name: cfg.GetString(fmt.Sprintf("check.name%d", i)),
				})
			}
		}
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	var wg = &sync.WaitGroup{}
	output := make([]string, 0)
	for _, row := range res {
		wg.Add(1)
		go func(row *UrlInfo) {
			defer wg.Done()
			resp, statusCode, err := request.HTTPRequest(request.POST, row.Url, strings.NewReader("{}"), &header)
			var s string
			if statusCode != 200 || err != nil {
				if err != nil {
					s = err.Error()
				}
				output = append(output, fmt.Sprintf("%s：statusCode:%d,err:%s,resp:%s", row.Name, statusCode, s, resp))
			}
		}(row)
	}
	wg.Wait()
	return output
}

// SystemInfo 报告系统消息
func SystemInfo(cfg *viper.Viper) error {
	percent := cfg.GetFloat64("app.percent") // 警告百分比0-95
	if percent > 95 {
		percent = 95
	}
	messageList := GetAllReportInfo(percent)
	if len(messageList) > 0 {
		err := SendTextMessage(cfg, messageList)
		if err != nil {
			return err
		}
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
