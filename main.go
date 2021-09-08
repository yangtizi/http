package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/yangtizi/http/fast"
	"github.com/yangtizi/http/post"
)

func main() {
	// fasthttpStartDemo()
	postDemo()
	// post.Default.URL("http://127.0.0.1:1234/ping").AddHeader("time", time.Now().String())

}

func postDemo() {

	post.Default.URL(
		"https://www.qq.com/help/player/create").TLS(
		"CNY.pem", "CNY.key").SetHeader(
		"X_ENTITY_KEY", "12897389719283")
	resp, err := post.Default.String("")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("返回结果", string(resp))
}

// 启动fast服务器
func fasthttpStartDemo() {
	fast.Register("/fast/Demo", onFastDemo)
	fast.StartServer(":8080")
}

func onFastDemo(ctx *fasthttp.RequestCtx) {
}
