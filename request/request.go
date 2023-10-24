package request

import (
	"io"
	"io/ioutil"
	"net/http"
)

const (
	POST = "POST"
	GET  = "GET"
)

// HTTPRequest 发起HTTP请求并返回响应内容  POST or GET
func HTTPRequest(method, urlStr string, params io.Reader, headers *map[string]string) (string, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, params)
	if err != nil {
		return "", 0, err
	}
	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}
	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", resp.StatusCode, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	return string(body), resp.StatusCode, nil
}
