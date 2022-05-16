package hcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
)

type LDEV struct {
	LdevID             int    `json:"ldevId"`
	ClprID             int    `json:"clprId"`
	EmulationType      string `json:"emulationType"`
	ByteFormatCapacity string `json:"byteFormatCapacity"`
	BlockCapacity      int    `json:"blockCapacity"`
	NumOfPorts         int    `json:"numOfPorts"`
	Ports              []struct {
		PortID          string `json:"portId"`
		HostGroupNumber int    `json:"hostGroupNumber"`
		HostGroupName   string `json:"hostGroupName"`
		Lun             int    `json:"lun"`
	} `json:"ports"`
	Attributes              []string `json:"attributes"`
	Label                   string   `json:"label"`
	Status                  string   `json:"status"`
	MpBladeID               int      `json:"mpBladeId"`
	Ssid                    string   `json:"ssid"`
	PoolID                  int      `json:"poolId"`
	NumOfUsedBlock          int      `json:"numOfUsedBlock"`
	IsFullAllocationEnabled bool     `json:"isFullAllocationEnabled"`
	ResourceGroupID         int      `json:"resourceGroupId"`
	DataReductionStatus     string   `json:"dataReductionStatus"`
	DataReductionMode       string   `json:"dataReductionMode"`
	IsAluaEnabled           bool     `json:"isAluaEnabled"`
	NaaId                   string   `json:"naaId"`
}

func GetAllLDEVs(session Session) []LDEV {
	resp, err := session.Request("GET", "/ldevs", nil, nil, nil)
	if err != nil {
		log.Printf("GetAllLDEVs error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	var ldevs []LDEV
	for _, item := range gjson.Get(string(byteBody), "data").Array() {
		tmpLdev := LDEV{}
		json.Unmarshal([]byte(item.String()), &tmpLdev)
		ldevs = append(ldevs, tmpLdev)
	}
	return ldevs
}

// 根据LdevId，获取LDEV的naaID；
func (l *LDEV) GetLdevNaaID(session Session) string {
	resp, err := session.Request("GET", "/ldevs/"+fmt.Sprintf("%d", l.LdevID), nil, nil, nil)
	if err != nil {
		log.Printf("GetLdevNaaID error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	l.NaaId = gjson.Get(string(byteBody), "naaId").String()
	return l.NaaId
}

// 根据十进制LdevId，获取LDEV的详细信息，包含NaaID；
func GetSpecifyLDEV(session Session, LdevID int64) LDEV {
	resp, err := session.Request("GET", "/ldevs/"+fmt.Sprintf("%d", LdevID), nil, nil, nil)
	if err != nil {
		log.Printf("GetLdevNaaID error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	targetLdev := LDEV{}
	json.Unmarshal([]byte(byteBody), &targetLdev)
	return targetLdev
}
