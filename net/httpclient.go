package net

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ReqConf 请求配置
type ReqConf struct {
	ReqURL string                 // 请求地址
	Method string                 // 请求方式，GET/POST
	Data   string                 // 提交数据
	Header http.Header            // 请求头
	Param  map[string]interface{} // 额外的参数
}

// HTTPClient 客户端
type HTTPClient struct {
	*ReqConf
	resp *http.Response
}

const (
	// TimeOutKey 超时设置key
	TimeOutKey = "timeOut"
	// 默认超时时间
	defaultTimeOut = time.Duration(5)
	// TotalTryTimesKey 重试次数key
	TotalTryTimesKey = "totalTryTimes"
	// 默认重试次数
	defaultTotalTryTimes = 3
)

// NewHTTPClient 创建一个客户端
func NewHTTPClient(req *ReqConf) *HTTPClient {
	return &HTTPClient{req, nil}
}

//Send 发送请求
/**
@return resp	响应体
@return err		错误信息
*/
func (c *HTTPClient) Send() (err error) {
	if 0 == len(c.ReqURL) {
		return errors.New("url cannot empty")
	}

	timeOut, ok := c.Param[TimeOutKey]
	if !ok {
		timeOut = defaultTimeOut
	}

	totalTryTime, ok := c.Param[TotalTryTimesKey]
	if !ok {
		totalTryTime = defaultTotalTryTimes
	}

	req, err := http.NewRequest(c.Method, c.ReqURL, strings.NewReader(c.Data))

	if nil != err {
		return err
	}

	// 设置header
	if nil != c.Header {
		for key, val := range c.Header {
			req.Header.Set(key, val[0])
		}
	}
	// 创建客户端
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("whttp.HttpReq found 302")
			return nil
		},
		Timeout: timeOut.(time.Duration) * time.Second,
	}

	// 发送请求，如果超时可尝试多次(依据配置)
	tryTimes := 1
	for tryTimes <= totalTryTime.(int) {
		c.resp, err = client.Do(req)
		if nil == err {
			break
		} else {
			if e, ok := err.(*url.Error); ok && e.Timeout() {
				log.Printf("whttp.HttpRep(%s) timeout, try %dth", c.ReqURL, tryTimes)
				tryTimes++
			} else {
				return err
			}
		}
	}

	return
}

// Close 关闭响应体
func (c *HTTPClient) Close() error {
	return c.resp.Body.Close()
}

// RecvBody 获取响应体
func (c *HTTPClient) RecvBody() *http.Response {
	return c.resp
}
