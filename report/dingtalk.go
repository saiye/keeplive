package report

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"game_go/request"
	"github.com/spf13/viper"
	"net/url"
	"strings"
	"sync"
	"time"
)

type SecretInfo struct {
	AccessToken string `json:"access_token"`
	Secret      string `json:"secret"`
}
type TextMessage struct {
	At      AtContent   `json:"at"`
	Text    TextContent `json:"text"`
	MsgType string      `json:"msgtype"`
}

type AtContent struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}
type TextContent struct {
	Content string `json:"content"`
}
type ResInfo struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func HmacSha256(key, data string) string {

	hash := hmac.New(sha256.New, []byte(key)) //创建对应的sha256哈希加密算法

	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum([]byte("")))
}

func (receiver *SecretInfo) MakeSign(timestamp int64) string {
	//把timestamp+"\n"+密钥当做签名字符串，使用HmacSHA256算法计算签名，然后进行Base64 encode，最后再把签名参数再进行urlEncode，得到最终的签名（需要使用UTF-8字符集）。
	inputString := fmt.Sprintf("%d\n%s", timestamp, receiver.Secret)
	h := hmac.New(sha256.New, []byte(receiver.Secret))
	h.Write([]byte(inputString))
	bas64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))
	urlEncodeStr := url.QueryEscape(bas64Str)
	return urlEncodeStr
}

func NewSecretInfo(accessToken string, secret string) *SecretInfo {
	return &SecretInfo{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

func (receiver *TextMessage) Send(config *viper.Viper) error {
	accessToken := config.GetString("dingtalk.access_token") // 读取配置
	secretVal := config.GetString("dingtalk.secret")         // 读取配置
	secret := NewSecretInfo(accessToken, secretVal)
	header := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	jsonV, err := json.Marshal(receiver)
	if err != nil {
		fmt.Errorf("json encode err:%v data: %v", err, *receiver)
		return err
	}
	readerData := bytes.NewReader(jsonV)
	timestamp := time.Now().UnixMilli()
	sign := secret.MakeSign(timestamp)
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", secret.AccessToken, timestamp, sign)

	respJson, statusCode, err := request.HTTPRequest(request.POST, url, readerData, &header)

	if respJson != "" {
		var resp ResInfo

		err2 := json.Unmarshal([]byte(respJson), &resp)
		if err2 != nil {
			return errors.New(fmt.Sprintf("debug:json Unmarshal err %v,jsonRes:%v", err2, respJson))
		}

		if resp.ErrCode != 0 || err != nil || statusCode != 200 {
			sendData := string(jsonV)
			return errors.New(fmt.Sprintf("debug:消息报告失败：resp %v,statusCode:%v ,参数 %v,url:%v", resp, statusCode, sendData, url))
		}
	}
	return err
}

// SendTextMessage 发送钉钉的text 消息
func SendTextMessage(cfg *viper.Viper, messageList []string) error {
	env := cfg.GetString("app.env") // 读取配置
	phoneList := GetPhoneList(cfg.GetString("dingtalk.phone_list"))
	keyword := cfg.GetString("dingtalk.keyword") // 读取警报关键词
	var wg = &sync.WaitGroup{}
	output := make([]string, 0)
	for _, msg := range messageList {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
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
			err := receiver.Send(cfg)
			if err != nil {
				output = append(output, err.Error())
			}
		}(msg)
	}
	wg.Wait()
	if len(output) > 0 {
		return errors.New(strings.Join(output, "\n"))
	}
	return nil
}
