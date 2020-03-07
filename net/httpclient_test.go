package net

import (
	"fmt"
	"net/http"
	"testing"
)

var c *HTTPClient

func TestMain(m *testing.M) {
	// 测试运行前
	c = NewHTTPClient(&ReqConf{ReqURL: "https://studygolang.com/pkgdoc", Method: "GET"})
	// 开始运行测试
	m.Run()
	// 测试运行结束后
	resp := c.RecvBody()
	fmt.Println(resp.Body.Read([]byte{}))
}

func TestSend(t *testing.T) {

	var resp *http.Response

	err := c.Send()
	if nil != err {
		t.Fatal(err)
	}
	defer func() {
		// 关闭响应体Body
		err = c.Close()
		if nil != err {
			t.Fatal(err)
		}
	}()

	resp = c.RecvBody()

	fmt.Println(resp.Header.Get("X-Request-Id"))
}
