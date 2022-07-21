package sannav

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

var (
	IP       string
	Username string
	Password string
	Schema   string = "https"
	Token    string
	tr       = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)

func init() {
	confByte, err := ioutil.ReadFile("./sannavcfg.json")
	if err != nil {
		log.Fatal("SANnav initialization failed: ", err)
	}
	IP = gjson.Get(string(confByte), "IP").String()
	Username = gjson.Get(string(confByte), "Username").String()
	Password = gjson.Get(string(confByte), "Password").String()
}

// 登录，获取访问token；
func Login() string {
	url := Schema + "://" + IP + "/external-api/v1/login/"
	log.Printf("request URL: %s\n", url)
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Basic YWRtaW5pc3RyYXRvcjoxcWF6QFdTWAo=")
	req.Header.Add("username", Username)
	req.Header.Add("password", Password)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("获取Token失败，状态码: %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	Token = gjson.Get(string(body), "sessionId").String()

	log.Printf("Token获取成功： %s\n", Token)

	return Token
}

func Logout() {
	url := Schema + "://" + IP + "/external-api/v1/login/"
	log.Printf("request URL: %s\n", url)
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Session "+Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Logout失败，状态码: %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	Token = gjson.Get(string(body), "sessionId").String()

	log.Printf("Token获取成功： %s\n", Token)
}
