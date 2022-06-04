package hcm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/tidwall/gjson"
)

// 基本Port结构体，其他端口类型继承自它；
type Port struct {
	PortID             string   `json:"portId"`
	PortType           string   `json:"portType"`
	PortAttributes     []string `json:"portAttributes"`
	PortSpeed          string   `json:"portSpeed"`
	LoopID             string   `json:"loopId"`
	FabricMode         bool     `json:"fabricMode"`
	LunSecuritySetting bool     `json:"lunSecuritySetting"`
	WWN                string   `json:"wwn"`
}

//判断slice中是否存在某个item
func IsExistItem(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

/*
	@title GetAllPorts
	@description 根据传递的参数获取存储Port；
	@auth 	Jason Xiong
	@param	 portType        string     端口类型："FIBRE", "SCSI", "ISCSI", "ENAS", "ESCON", "FICON", "FCoE", "HNASS", "HNASU"
	@param   portAttributes  string     "TAR", "MCU", "RCU", "ELUN"
	@param	 portId			 string		端口ID，例如: CL1-A
	@param   detailInfoType  string		可选参数为：logins
*/
func GetAllPorts(session *Session, portType, portAttributes, portId, detailInfoType string) []PortDetail {
	var Parameters = map[string]string{}
	PortTypes := []string{"FIBRE", "SCSI", "ISCSI", "ENAS", "ESCON", "FICON", "FCoE", "HNASS", "HNASU"}
	if portType != "" && portId != "" {
		log.Fatal("portType和portId参数同时只能指定一个")
	}

	if portType != "" && IsExistItem(portType, PortTypes) {
		Parameters["portType"] = portType
	}

	PortAttributes := []string{"TAR", "MCU", "RCU", "ELUN"}
	if portAttributes != "" && portId != "" {
		log.Fatal("portAttributes和portId参数同时只能指定一个")
	}

	if portAttributes != "" && IsExistItem(portAttributes, PortAttributes) {
		Parameters["portAttributes"] = portAttributes
	}

	// portId参数可以在：VSP Gx00,VSP G1000,VSP G1500,VSP Fx00, VSP F1500存储时指定；
	// 指定portId参数，需要确保指定了detailInfoType参数；
	// 指定portId参数，则不能同时指定portType 和 portAttributes参数；
	if portId != "" {
		Parameters["portId"] = portId
	}

	// detailInfoType: 取值：logins
	if detailInfoType != "" {
		Parameters["detailInfoType"] = detailInfoType
	}

	resp, _ := session.Request("GET", "/ports", Parameters, nil, nil)
	byteBody, _ := ioutil.ReadAll(resp.Body)
	var ports []PortDetail
	for _, port := range gjson.Get(string(byteBody), "data").Array() {
		tmpport := PortDetail{}
		err := json.Unmarshal([]byte(port.String()), &tmpport)
		if err != nil {
			log.Fatal("Unmarshal Port Failed.")
		}
		ports = append(ports, tmpport)
	}
	return ports
}

func GetSpecificPort(session *Session, portID string) PortDetail {
	resp, err := session.Request("GET", "/ports/"+portID, nil, nil, nil)
	if err != nil {
		log.Printf("GetPortDetail error:%s\n", err)
	}
	byteBody, _ := ioutil.ReadAll(resp.Body)
	retPort := PortDetail{}
	json.Unmarshal(byteBody, &retPort)
	return retPort
}

type PortLogin struct {
	LoginWwn    string `json:"loginWwn"`
	WwnNickName string `json:"wwnNickName"`
	IsLoggedIn  bool   `json:"isLoggedIn"`
	HostGroupID string `json:"hostGroupId,omitempty"`
}

type PortDetail struct {
	Port
	PortConnection string      `json:"portConnection"`
	Wwn            string      `json:"wwn"`
	Logins         []PortLogin `json:"logins"`
}

type FCoEPort struct {
	Port
	PortConnection      string `json:"portConnection"`
	Wwn                 string `json:"wwn"`
	StaticMacAddress    string `json:"staticMacAddress"`
	VLanID              string `json:"vLanId"`
	DynamicMacAddress   string `json:"dynamicMacAddress"`
	VirtualPortStatus   string `json:"virtualPortStatus"`
	VirtualPortID       string `json:"virtualPortId"`
	FcoeSwitchControlID string `json:"fcoeSwitchControlId"`
}

type TCPOption struct {
	Ipv6Mode         bool `json:"ipv6Mode"`
	SelectiveAckMode bool `json:"selectiveAckMode"`
	DelayedAckMode   bool `json:"delayedAckMode"`
	IsnsService      bool `json:"isnsService"`
	TagVLan          bool `json:"tagVLan"`
}

type ISCSIPort struct {
	Port
	Logins []struct {
		LoginIscsiName string `json:"loginIscsiName"`
	} `json:"logins"`
	TcpOption            TCPOption `json:"tcpOption"`
	TCPMtu               int       `json:"tcpMtu"`
	IscsiWindowSize      string    `json:"iscsiWindowSize"`
	KeepAliveTimer       int       `json:"keepAliveTimer"`
	TCPPort              string    `json:"tcpPort"`
	Ipv4Address          string    `json:"ipv4Address"`
	Ipv4Subnetmask       string    `json:"ipv4Subnetmask"`
	Ipv4GatewayAddress   string    `json:"ipv4GatewayAddress"`
	Ipv6LinkLocalAddress struct {
		Status         string `json:"status"`
		AddressingMode string `json:"addressingMode"`
		Address        string `json:"address"`
	} `json:"ipv6LinkLocalAddress"`
	Ipv6GlobalAddress struct {
		Status         string `json:"status"`
		AddressingMode string `json:"addressingMode"`
		Address        string `json:"address"`
	} `json:"ipv6GlobalAddress"`
	Ipv6GatewayGlobalAddress struct {
		Status         string `json:"status"`
		Address        string `json:"address"`
		CurrentAddress string `json:"currentAddress"`
	} `json:"ipv6GatewayGlobalAddress"`
}
