package fast

import (
	"sync"

	"github.com/valyala/fasthttp"
)

// ! import "github.com/yangtizi/go/fast"

var mpHandler sync.Map

// Register 注册回调
func Register(strPath string, cb func(*fasthttp.RequestCtx)) {
	mpHandler.Store(strPath, cb)
}
