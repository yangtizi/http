package fast

import (
	"github.com/valyala/fasthttp"
	"github.com/yangtizi/log/zaplog"
)

// ! import "github.com/yangtizi/http/fast"

// StartServer (地址)
func StartServer(strAddress string) {
	defer zaplog.Ins.Infof("服务器", strAddress, "已经关闭") //

	if len(strAddress) <= 0 { //
		zaplog.Ins.Errorf("错误的Address = [%s]", strAddress) //
		return
	} //
	zaplog.Ins.Infof("服务器 [%s] 正在开启", strAddress)               //
	err := fasthttp.ListenAndServe(strAddress, fastHTTPHandler) //
	if err != nil {                                             //
		zaplog.Ins.Errorf("fasthttp.ListenAndServe错误 = [%v]", err) //
	} //

}
