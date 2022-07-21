package hcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
)

type HostGroup struct {
	HostGroupID     string   `json:"hostGroupId"`
	PortID          string   `json:"portId"`
	HostGroupNumber int      `json:"hostGroupNumber"`
	HostGroupName   string   `json:"hostGroupName"`
	HostMode        string   `json:"hostMode"`
	HostModeOptions []int    `json:"hostModeOptions"`
	ResourceGroupID int      `json:"resourceGroupId"`
	IsDefined       bool     `json:"isDefined"`
	NumOfLun        int      `json:"numOfLun"`
	Luns            []LUN    `json:"luns"`
	NumOfHost       int      `json:"numOfHost"`
	HostWWNs        []string `json:"hosts"`
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
	resp, err := session.Request("GET", "/host-groups/"+hostgroup.HostGroupID, Parameters, nil, nil)
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
	hostgroup.Luns = luns
	hostgroup.NumOfLun = len(luns)
	return luns
}

func (hostgroup *HostGroup) GetHostgroupHosts(session *Session) {
	Parameters := map[string]string{}
	Parameters["portId"] = hostgroup.PortID
	Parameters["hostGroupNumber"] = fmt.Sprintf("%d", hostgroup.HostGroupNumber)
	resp, err := session.Request("GET", "/host-wwns", Parameters, nil, nil)
	if err != nil {
		log.Printf("GetHostgroupHosts error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	for _, tmpHost := range gjson.Get(string(byteBody), "data").Array() {
		hostgroup.HostWWNs = append(hostgroup.HostWWNs, gjson.Get(tmpHost.String(), "hostWwn").String())
	}
	hostgroup.NumOfHost = len(hostgroup.HostWWNs)
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

func SetNickName(session *Session, portId, hostWwn, nickname string, hostGroupNumber int) {
	if nickname == "" {
		log.Printf("别名为空，表示删除该wwn号的别名.")
	}
	reqBody := map[string]string{}
	reqBody["wwnNickname"] = nickname
	// byteReqBody, _ := json.Marshal(reqBody)

	resp, err := session.Request("PUT", "/host-wwns/"+portId+","+fmt.Sprintf("%d", hostGroupNumber)+","+hostWwn, nil, reqBody, nil)
	if err != nil {
		log.Printf("Set Nickname error:%s\n", err)
		return
	}
	if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 202 {
		log.Printf("别名设置成功: %s\n", nickname)
	} else {
		byteBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteBody))
		fmt.Println(resp.StatusCode)
	}
}

func GetAllHostWWNs(session *Session, portID string) ([]HostWWN, error) {
	Parameters := map[string]string{}
	Parameters["portId"] = portID
	resp, err := session.Request("GET", "/host-wwns", Parameters, nil, nil)
	if err != nil {
		log.Printf("Get port %s HostWWNs error:%s\n", portID, err)
		return nil, err
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	hosts := make([]HostWWN, 20)
	for _, item := range gjson.Get(string(byteBody), "data").Array() {
		host := HostWWN{}
		json.Unmarshal([]byte(item.String()), &host)
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func GetSpecificHostWWN(session *Session, portId, hostGroupNumber, hostWwn string) (HostWWN, error) {
	resp, err := session.Request("GET", "/host-wwns/"+portId+","+hostGroupNumber+","+hostWwn, nil, nil, nil)
	if err != nil {
		log.Printf("Get Specific Host WWN error:%s\n", err)
		return HostWWN{}, err
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	host := HostWWN{}
	json.Unmarshal(byteBody, &host)
	return host, nil
}

func RegisterWWNIntoHostgroup(session *Session, portId, hostWwn string, hostGroupNumber int) {
	resp, err := session.Request("POST", "/host-wwns/"+portId+","+fmt.Sprintf("%d", hostGroupNumber)+","+hostWwn, nil, nil, nil)
	if err != nil {
		log.Printf("Register WWN into Hostgroup error:%s\n", err)
		return
	}
	// byteBody, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()
	// hostwwn := HostWWN{}
	// json.Unmarshal(byteBody, &hostwwn)
	if resp.StatusCode != 200 {
		log.Printf("添加WWN:%s 到主机组:%s 失败;\n", hostWwn, portId+","+fmt.Sprintf("%d", hostGroupNumber))
		return
	}
	log.Printf("Register %s into %s success;\n", hostWwn, portId+","+fmt.Sprintf("%d", hostGroupNumber))
}
