package hcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
)

type HostGroup struct {
	HostGroupID     string `json:"hostGroupId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostGroupName   string `json:"hostGroupName"`
	HostMode        string `json:"hostMode"`
	HostModeOptions []int  `json:"hostModeOptions"`
	ResourceGroupID int    `json:"resourceGroupId"`
	IsDefined       bool   `json:"isDefined"`
}

func GetAllHostgroups(session *Session) []HostGroup {
	var hostgroups []HostGroup
	resp, err := session.Request("GET", "/host-groups", nil, nil, nil)
	if err != nil {
		log.Printf("GetAllHostgroups error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	for _, hostgroup := range gjson.Get(string(byteBody), "data").Array() {
		tmpHostgroup := HostGroup{}
		json.Unmarshal([]byte(hostgroup.String()), &tmpHostgroup)
		hostgroups = append(hostgroups, tmpHostgroup)
	}
	return hostgroups
}

func (hostgroup *HostGroup) GetHostgroupDetail(session *Session) {
	Parameters := map[string]string{}
	Parameters["detailInfoType"] = "resourceGroup"
	resp, err := session.Request("GET", "/host-groups"+hostgroup.HostGroupID, Parameters, nil, nil)
	if err != nil {
		log.Printf("GetHostgroupDetail error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	tmpHostgroup := HostGroup{}
	json.Unmarshal(byteBody, &tmpHostgroup)
	hostgroup = &tmpHostgroup
}

func (hostgroup *HostGroup) GetHostgroupLuns(session *Session) []LUN {
	Parameters := map[string]string{}
	Parameters["portId"] = hostgroup.PortID
	Parameters["hostGroupNumber"] = fmt.Sprintf("%d", hostgroup.HostGroupNumber)
	resp, err := session.Request("GET", "/luns", Parameters, nil, nil)
	if err != nil {
		log.Printf("GetHostgroupLuns error:%s\n", err)
	}

	var luns []LUN
	byteBody, _ := ioutil.ReadAll(resp.Body)
	for _, item := range gjson.Get(string(byteBody), "data").Array() {
		lun := LUN{}
		json.Unmarshal([]byte(item.String()), &lun)
		luns = append(luns, lun)
	}
	return luns
}

func (hostgroup *HostGroup) GetHostgroupHosts(session *Session) {

}

type ISCSITarget struct {
	HostGroupID          string `json:"hostGroupId"`
	PortID               string `json:"portId"`
	HostGroupNumber      int    `json:"hostGroupNumber"`
	HostGroupName        string `json:"hostGroupName"`
	IscsiName            string `json:"iscsiName"`
	AuthenticationMode   string `json:"authenticationMode"`
	IscsiTargetDirection string `json:"iscsiTargetDirection"`
	HostMode             string `json:"hostMode"`
	HostModeOptions      []int  `json:"hostModeOptions"`
}

type HostMode struct {
	HostModeID      int    `json:"hostModeId"`
	HostModeName    string `json:"hostModeName"`
	HostModeDisplay string `json:"hostModeDisplay"`
}

type HostModeOption struct {
	HostModeOptionID          int    `json:"hostModeOptionId"`
	HostModeOptionDescription string `json:"hostModeOptionDescription"`
}

type HostWWN struct {
	HostWwnID       string `json:"hostWwnId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostGroupName   string `json:"hostGroupName"`
	HostWwn         string `json:"hostWwn"`
	WwnNickname     string `json:"wwnNickname"`
}

type HostISCSI struct {
	HostIscsiID     string `json:"hostIscsiId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostGroupName   string `json:"hostGroupName"`
	IscsiName       string `json:"iscsiName"`
	IscsiNickname   string `json:"iscsiNickname"`
}
