package get

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/yangtizi/log/zaplog"
)

// Bytes 函数
func Bytes(strURL string) ([]byte, error) {
	resp, err := http.Get(strURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	return body, nil
}

func MyGet(strURL string, rsp interface{}) error {
	resp, err := http.Get(strURL)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zaplog.Ins.Errorf("%s", err.Error())
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = json.Unmarshal(body, rsp)
	return err
}
