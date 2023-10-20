package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	type args struct {
		method  string
		urlStr  string
		params  url.Values
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "post-test",
			args: args{
				method: "POST",
				urlStr: "https://api.mch.weixin.qq.com/v3/pay/transactions/app",
				params: url.Values{
					"appid": {"111111111"},
				},
				headers: map[string]string{
					"Content-Type": "application/json",
					"Accept":       "application/json",
				},
			},
			want:    "machid",
			wantErr: false,
		},
	}
	type Info struct {
		Code   string `json:"code"`
		Detail struct {
			Location string `json:"location"`
			Value    string `json:"value"`
		} `json:"detail"`
		Message string `json:"message"`
	}
	var info Info
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HTTPRequest(tt.args.method, tt.args.urlStr, tt.args.params, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err2 := json.Unmarshal([]byte(got), &info)
			if err2 != nil {
				t.Errorf("json Unmarshal err %v", err2)
				return
			}
			fmt.Println("info.Message:", info.Message)
			if !strings.Contains(info.Message, tt.want) {
				t.Errorf("HTTPRequest() got = %v, want %v", info.Message, tt.want)
			}
		})
	}
}
