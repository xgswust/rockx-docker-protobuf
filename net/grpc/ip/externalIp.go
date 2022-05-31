package ip

import (
	"io/ioutil"
	"net/http"
)

func GetExternal() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		println("error", err.Error())
	}
	defer resp.Body.Close()
	return parseResponseString(resp)
}

func parseResponseString(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body) // response.Body 是一个数据流
	return string(body), err                   // 将 io数据流转换为string类型返回
}
