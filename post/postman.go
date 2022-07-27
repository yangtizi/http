package post

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/yangtizi/log/zaplog"
)

type TPostMan struct {
	client    *http.Client
	strURL    string
	mapHead   http.Header
	req       *http.Request
	transport *http.Transport
	lock      sync.Mutex
}

// 创建一个
func NewPostMan() *TPostMan {
	return &TPostMan{
		client:  &http.Client{},
		mapHead: make(http.Header),
	}
}

var Default = NewPostMan()

// 关闭. 通常在defer 里面调用,  用来清理一些无用的玩意
func (me *TPostMan) Close() {
	// 清理
	if me.client != nil {
		me.client.CloseIdleConnections()
	}
}

// 上锁用,
func (me *TPostMan) Lock() {
	me.lock.Lock()
}
func (me *TPostMan) Unlock() {
	me.lock.Unlock()
}

func (me *TPostMan) SetURL(strURL string) string {
	me.strURL = strURL
	return me.strURL
}
func (me *TPostMan) GetURL() string {
	return me.strURL
}

func (me *TPostMan) URL(strURL string) *TPostMan {
	me.SetURL(strURL)
	return me
}

func (me *TPostMan) AddHeader(strKey, strValue string) *TPostMan {
	me.mapHead.Add(strKey, strValue)
	return me
}

func (me *TPostMan) SetHeader(strKey, strValue string) *TPostMan {
	me.mapHead.Set(strKey, strValue)
	return me
}

// 缩减代码
func (me *TPostMan) URLEncodedHeader() *TPostMan {
	me.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return me
}

func (me *TPostMan) String(strContent string) ([]byte, error) {
	req, err := http.NewRequest("POST", me.strURL, strings.NewReader(strContent))

	if err != nil {
		return nil, err
	}

	req.Header = me.mapHead
	me.req = req

	return me.do()
}

func (me *TPostMan) TLS(strCertFile, strKeyFile string) *TPostMan {
	cliCrt, err := tls.LoadX509KeyPair(strCertFile, strKeyFile)

	if err != nil {
		zaplog.Ins.Errorf("TLS(strCertFile=[%s], strKeyFile=[%s]", strCertFile, strKeyFile)
		return me
	}

	if me.transport == nil {
		me.transport = &http.Transport{}
	}

	if me.transport.TLSClientConfig == nil {
		me.transport.TLSClientConfig = &tls.Config{}
	}

	// 添加进去
	me.transport.TLSClientConfig.Certificates = append(me.transport.TLSClientConfig.Certificates, cliCrt)
	return me
}

func (me *TPostMan) do() ([]byte, error) {
	if me.req == nil {
		return nil, errors.New("no request")
	}

	me.client.Transport = me.transport

	rsp, err := me.client.Do(me.req)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	return body, err
}
