package netapp

import (
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

var (
	BaseURL    string
	UserName   string
	Password   string
	BasicToken string
)

func init() {
	confByte, err := ioutil.ReadFile("netappcfg.json")
	if err != nil {
		log.Printf("加载配置文件错误: %v, 退出...\n", err)
		os.Exit(1)
	}
	BaseURL = gjson.Get(string(confByte), "baseurl").String()
	UserName = gjson.Get(string(confByte), "username").String()
	Password = gjson.Get(string(confByte), "password").String()
	//将用户名和密码进行Base64编码，用来进行Basic认证
	BasicToken = base64.StdEncoding.EncodeToString([]byte(UserName + ":" + Password))

	// BaseURL = "https://192.204.2.6"
	// UserName = "readuser"
	// Password = "Crc8@read"
	// BasicToken = base64.StdEncoding.EncodeToString([]byte(UserName + ":" + Password))
}

func ApiGetRequest(url_postfix string) []byte {
	url := BaseURL + url_postfix
	// 忽略https证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", "Basic "+BasicToken)
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("During request %v Error: %v\n", url_postfix, err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	return body
}
