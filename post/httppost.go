package post

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/yangtizi/log/zaplog"
	"golang.org/x/net/proxy"
)

// JSON (strURL 网址,  strJson POST内容)
func JSON(strURL, strJSON string) string {
	resp, err := http.Post(
		strURL,
		"application/json",
		strings.NewReader(strJSON))

	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
		return string(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
		return string(err.Error())
	}

	// log.Println(string(body))
	return string(body)
}

// Bytes 这里是发送BUFF内容
func Bytes(strURL string, buf []byte) ([]byte, error) {
	resp, err := http.Post(
		strURL,
		"",
		bytes.NewReader(buf))

	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
	}

	// log.Println(string(body))PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv
	return body, err
}

// HTTPProxy 使用HTTP代理
func HTTPProxy(strURL string, buf []byte, strProxy string, strUser string, strPassword string) ([]byte, error) {

	// 配置代理
	urli := url.URL{}
	urlproxy, _ := urli.Parse(strProxy)
	// 拿出代理
	client := &http.Client{}

	if len(strProxy) > 7 {
		zaplog.Ins.Infof("使用代理 = [%v]", urlproxy)
		client.Transport = &http.Transport{Proxy: http.ProxyURL(urlproxy)}
	} else {
		zaplog.Ins.Infof("不使用代理 %s", strProxy)
	}

	// 请求
	req, err := http.NewRequest("POST", strURL, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	// 设置头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-type", "application/octet-stream")
	req.Header.Set("User-Agent", "MicroMessenger Client")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()
	defer resp.Body.Close()

	// 读取 respon
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Sock5Proxy Sock5代理
func Sock5Proxy(strURL string, buf []byte, strProxy string, strUser string, strPassword string) ([]byte, error) {
	auth := &proxy.Auth{
		User:     strUser,
		Password: strPassword,
	}

	if strURL == "" && strPassword == "" {
		auth = nil
	}

	dialer, err := proxy.SOCKS5("tcp", strProxy, auth, proxy.Direct)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	if len(strProxy) > 7 {
		zaplog.Ins.Infof("使用代理%s", strProxy)
		client.Transport = &http.Transport{Dial: dialer.Dial}
	} else {
		zaplog.Ins.Infof("不用代理%s", strProxy)
	}

	// 请求
	req, err := http.NewRequest("POST", strURL, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	// 设置头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-type", "application/octet-stream")
	req.Header.Set("User-Agent", "MicroMessenger Client")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()
	defer resp.Body.Close()

	// 读取 respon
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Https 地址,
func Https(strURL, strCertFile, strKeyFile string) ([]byte, error) {
	// 具体的证书加载对象
	cliCrt, err := tls.LoadX509KeyPair(strCertFile, strKeyFile)
	if err != nil {
		return nil, err
	}

	// 把上面的准备内容传入 client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
	}

	req, err := http.NewRequest("POST", strURL, strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()
	defer resp.Body.Close()

	// 读取 respon
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func MyPost(strURL string, req interface{}, rsp interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	payload := bytes.NewReader(b)

	r, err := http.NewRequest("POST", strURL, payload)
	if err != nil {
		return err
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rsp)
	return err
}
