package report

import (
	"errors"
	"fmt"
	"game_go/system"
	"github.com/spf13/viper"
	"strings"
)

var _config *viper.Viper

func GetConfig(configDir string) (*viper.Viper, error) {
	if _config == nil {
		cfg, err := system.NewConfig(configDir, "keeplive", "ini")
		if err != nil {
			return nil, err
		}
		_config = cfg
	}
	return _config, nil
}

// ReportMessage 报告系统状态
func ReportMessage(configDir string) error {
	cfg, err := GetConfig(configDir)
	if err != nil {
		return err
	}
	env := cfg.GetString("app.env")         // 读取配置
	keyword := cfg.GetString("app.keyword") // 读取警报关键词
	messageList := GetAllReportInfo(1)
	phoneList := GetPhoneList(cfg.GetString("dingtalk.phone_list"))
	message := ""
	for _, msg := range messageList {
		receiver := &TextMessage{
			At: AtContent{
				AtMobiles: phoneList,
				AtUserIds: nil,
				IsAtAll:   len(phoneList) == 0,
			},
			Text: TextContent{
				Content: fmt.Sprintf("环境:%s,message:%s:%s", env, keyword, msg),
			},
			MsgType: "text",
		}
		res := receiver.Send(cfg)
		if res != nil {
			messag += res.Error()
		}
	}
	if messag != "" {
		return errors.New(messag)
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
