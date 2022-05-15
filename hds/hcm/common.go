package hcm

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const (
	SCHEMA = "http" // or "http"
	HOST   = "192.204.1.91"
	PORT   = 23450 // or 23450
)

// 默认忽略https证书；
var tr = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func GenerateToken(deviceID, username, password string) (string, int64) {
	basicURL := URL(deviceID)
	URL := basicURL + "/sessions/"

	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		panic(err)
	}
	reqest.Header.Add("Accept", "application/json")
	reqest.Header.Add("Content-Type", "application/json")
	reqest.Header.Add("Authorization", "Basic "+EncodeCredentials(username, password))
	// reqest.Header.Add("Authorization", "Basic aGlhYTpQQHNzdzByZA==")

	response, err := client.Do(reqest)
	if err != nil {
		log.Printf("Generate Token error: %v\n", err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	Token := gjson.Get(string(body), "token").String()
	SessionID := gjson.Get(string(body), "sessionId").Int()
	return Token, SessionID
}

func URL(deviceID string) string {
	return SCHEMA + "://" + HOST + ":" + fmt.Sprintf("%d", PORT) + "/ConfigurationManager/v1/objects/storages/" + deviceID
}

type Session struct {
	SessionID        int64     `json:"sessionId"`
	UserID           string    `json:"userId"`
	IPAddress        string    `json:"ipAddress"`
	CreatedTime      time.Time `json:"createdTime"`
	LastAccessedTime time.Time `json:"lastAccessedTime"`
	http             http.Client
	DeviceID         string
	Token            string
}

func NewSession(deviceID, username, password string) (*Session, error) {
	if deviceID == "" || username == "" || password == "" {
		return nil, errors.New("deviceID,username and password cannot be empty")
	}

	token, sessionId := GenerateToken(deviceID, username, password)
	session := &Session{
		DeviceID:  deviceID,
		Token:     token,
		SessionID: sessionId,
	}

	return session, nil
}

//CloseSession purpose is to end the session。
func (session *Session) CloseSession() (err error) {
	_, err = session.Request("DELETE", "/sessions/"+fmt.Sprintf("%d", session.SessionID), nil, nil, nil)
	// fmt.Println(body)
	return err
}

func (session *Session) Request(method string, URI string, Parameters map[string]string, body, resp interface{}) (*http.Response, error) {
	if method == "" || URI == "" {
		return nil, errors.New("missing Method or URI")
	}

	endpoint := URL(session.DeviceID) + URI
	log.Printf("session Request endpoint:%s\n", endpoint)
	// create a http Request pointer
	var req *http.Request

	if body != nil {
		// Parse out body struct into JSON
		bodyBytes, _ := json.Marshal(body)

		// Create a new request
		req, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(bodyBytes))
	} else {
		// Create a new request
		req, _ = http.NewRequest(method, endpoint, nil)
	}

	// Add the mandatory headers to the request
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "session "+session.Token)

	// Create an URI query object
	if len(Parameters) > 0 {
		a := req.URL.Query()
		for k, v := range Parameters {
			a.Add(k, v)
		}
		req.URL.RawQuery = a.Encode()
	}

	// Perform request
	httpResp, err := session.http.Do(req)
	if err != nil {
		return httpResp, err
	}
	httpResp.Body.Close()

	return httpResp, nil
}
