package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

// Data 数据类型
type Data struct {
	ID string
	// 响应体
	Val RespResult
}

// FinalResult 最终整合的结果
/* type FinalResult struct {
	CatVal []Body
	DogVal []Body
} */

// RespResult 响应信息
type RespResult struct {
	Code int
	Msg  string
	Data []Body
}

// Body 响应body
type Body struct {
	ID      int
	Content string
}

var wg sync.WaitGroup
var errReq = errors.New("请求失败")

// 并发操作处理，最后汇总结果
func main() {
	var (
		data        = make(chan Data, 2)
		call        = map[string]string{"cat": "cat-url", "dog": "dog-url"}
		resultSlice []Data
		//result FinalResult
		result = make(map[string][]Body)
	)

	for k, v := range call {
		//wg.Add(1)
		go func(k, url string) {
			doReq(k, url, data)
		}(k, v)
	}

	//wg.Wait()
	//fmt.Println(reflect.TypeOf(errReq))

	for range call {
		result := <-data
		resultSlice = append(resultSlice, result)
	}
	// 根据响应结果的数据结构具体解析
	for _, v := range resultSlice {
		fmt.Println(reflect.ValueOf(v))

		//if v.Val.Code == 0 {
		result[v.ID] = v.Val.Data
		//}
	}

	fmt.Println(result)

	resBytes, err := json.Marshal(result)
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(string(resBytes))
}

func doReq(k, url string, data chan Data) {

	fmt.Printf("向%s发送请求\n", url)

	start := time.Now()
	fake := rand.Intn(5)
	time.Sleep(time.Duration(fake) * time.Second)
	end := time.Since(start)

	fmt.Printf("%d秒后得到响应结果\n", end)

	sucOrFail := []bool{true, true}

	data <- getBody(k, sucOrFail[rand.Intn(2)])

	//wg.Done()
}

// 模拟响应结果
func getBody(k string, fail bool) Data {

	var val RespResult

	if fail {
		val = RespResult{
			Code: -1,
			Msg:  errReq.Error(),
		}
	} else {
		val = RespResult{
			Code: 0,
			Msg:  "success",
			Data: []Body{
				{1, fmt.Sprintf("%s 's value 1", k)},
				{2, fmt.Sprintf("%s 's value 2", k)},
			},
		}
	}

	return Data{k, val}
}
