package request

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	type args struct {
		method  string
		urlStr  string
		params  map[string]interface{}
		headers map[string]string
	}
	tests := []struct {
		name           string
		args           args
		want           string
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "post-test",
			args: args{
				method: POST,
				urlStr: "https://api.mch.weixin.qq.com/v3/pay/transactions/app",
				params: map[string]interface{}{
					"appid": "111111111",
				},
				headers: map[string]string{
					"Content-Type": "application/json",
					"Accept":       "application/json",
				},
			},
			want:           "mchid",
			wantStatusCode: 400,
			wantErr:        false,
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
	var readerData io.Reader
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.params != nil {
				jsonV, err := json.Marshal(tt.args.params)
				if err != nil {
					t.Errorf("Json error = %v", err)
					return
				}
				readerData = bytes.NewReader(jsonV)
			}
			got, statusCode, err := HTTPRequest(tt.args.method, tt.args.urlStr, readerData, &tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if statusCode != tt.wantStatusCode {
				t.Errorf("HTTPRequest() StatusCode = %v, wantStatusCode %v", statusCode, tt.wantStatusCode)
				return
			}
			err2 := json.Unmarshal([]byte(got), &info)
			if err2 != nil {
				t.Errorf("json Unmarshal err %v", err2)
				return
			}
			if !strings.Contains(info.Message, tt.want) {
				t.Errorf("HTTPRequest() got = %v, want %v", info.Message, tt.want)
				return
			}
		})
	}
}
