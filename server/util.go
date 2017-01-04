package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TikiConfig struct {
	IP      string
	Port    uint
	DB      uint
	Timeout uint
	API     string
}

var cfg *TikiConfig

func SetConfig(c *TikiConfig) {
	cfg = c
}

/* POST 方式
 * @param url 目标地址
 * @param params POST 参数
 * @return response body
 */
func PostAPI(url string, params map[string]interface{}) ([]byte, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(paramsFormator(params)))
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return []byte(body), nil
}

// Post JSON to API
/*
 * @param url string 目标地址
 * @param json string 提交的json
 * @return response body
 */
func PostJSON(url string, json string) ([]byte, error) {
	resp, err := http.Post(url, "application/json",
		strings.NewReader(json))
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return []byte(body), nil
}

/* GET 方式
 * @param url 目标地址
 * @param params 参数
 * @return response body
 */
func GetAPI(url string, params map[string]interface{}) ([]byte, error) {
	resp, err := http.Get(url + "?" + paramsFormator(params))
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return []byte(body), nil
}

// 构造参数 (URLEncode)
func paramsFormator(params map[string]interface{}) string {
	if params == nil || len(params) == 0 { // 无参数
		return ""
	}
	var result string
	for key, val := range params {
		result += fmt.Sprintf("%s=%v&", key, val)
	}
	return strings.TrimRight(result, "&")
}
