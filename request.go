package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPRequest 发起HTTP请求并返回响应内容
func HTTPRequest(method, urlStr string, params url.Values, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, urlStr, strings.NewReader(params.Encode()))
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
