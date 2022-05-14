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
	"os"
	"reflect"
	"time"

	"github.com/tidwall/gjson"
)

const (
	SCHEMA = "https" // or "http"
	HOST   = "192.204.1.91"
	PORT   = 23451 // or 23450
)

// var (
// 	SN       string
// 	UserName string
// 	Password string
// )

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

func GetAllStorages() []StorageSystem {
	var storages []StorageSystem
	url := URL("")
	log.Printf("GetAllStorages:%s\n", url)

	client := &http.Client{Transport: tr}
	reqest, _ := http.NewRequest("GET", url, nil)

	response, _ := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("获取存储列表失败: %v\n", err)
		os.Exit(1)
	}

	for _, strStorage := range gjson.Get(string(body), "data").Array() {
		storage := StorageSystem{}
		json.Unmarshal([]byte(strStorage.String()), &storage)
		storages = append(storages, storage)
	}
	return storages
}

func GetDeviceIDBySN(sn int) string {
	var deviceID string
	storages := GetAllStorages()
	for _, storage := range storages {
		if storage.SerialNumber == sn {
			deviceID = storage.StorageDeviceID
		}
	}
	return deviceID
}

func (session *Session) Request(method string, URI string, Parameters, body, resp interface{}) error {
	if method == "" || URI == "" {
		return errors.New("missing Method or URI")
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

	if method != "DELETE" {
		// Create an URI query object
		a := req.URL.Query()

		// Add the query parameters to the URI query object
		t := reflect.TypeOf(Parameters)
		v := reflect.ValueOf(Parameters)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name, fmt.Sprintf("%v", v.Field(i).Interface()))
			a.Add(t.Field(i).Name, fmt.Sprintf("%v", v.Field(i).Interface()))
		}
	}

	// Perform request
	httpResp, err := session.http.Do(req)
	if err != nil {
		return err
	}

	// Cleanup Response
	defer httpResp.Body.Close()

	switch httpResp.StatusCode {
	case 200, 201, 202:
		// Decode JSON of response into our interface defined for the specific request sent
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}

		// Unmarshal the body into a struct
		bodyByte := json.Unmarshal(body, resp)

		return bodyByte
	case 204:
		return nil

	case 422:
		return fmt.Errorf("HTTP status codes: %d, detail: %v", httpResp.StatusCode, httpResp.Body)

	default:
		return fmt.Errorf("HTTP status codes: %d", httpResp.StatusCode)
	}
}

//CloseSession purpose is to end the session。
func (session *Session) CloseSession() (err error) {
	err = session.Request("DELETE", "/sessions/"+fmt.Sprintf("%d", session.SessionID), nil, nil, nil)
	return err
}

type StorageSystem struct {
	StorageDeviceID        string `json:"storageDeviceId"`
	Model                  string `json:"model"`
	SerialNumber           int    `json:"serialNumber"`
	SvpIP                  string `json:"svpIp"`
	MappWebServerHTTPSPort int    `json:"mappWebServerHttpsPort"`
	RmiPort                int    `json:"rmiPort"`
	Ctl1IP                 string `json:"ctl1Ip"`
	Ctl2IP                 string `json:"ctl2Ip"`
	DkcMicroVersion        string `json:"dkcMicroVersion"`
	CommunicationModes     []struct {
		CommunicationMode string `json:"communicationMode"`
	} `json:"communicationModes"`
	IsSecure              bool   `json:"isSecure"`
	LanConnectionProtocol string `json:"lanConnectionProtocol"`
	TargetCtl             string `json:"targetCtl"`
	UsesSvp               bool   `json:"usesSvp"`
}

type Job struct {
	JobID         int       `json:"jobId"`
	Self          string    `json:"self"`
	UserID        string    `json:"userId"`
	Status        string    `json:"status"`
	State         string    `json:"state"`
	CreatedTime   time.Time `json:"createdTime"`
	UpdatedTime   time.Time `json:"updatedTime"`
	CompletedTime time.Time `json:"completedTime"`
	Request       struct {
		RequestURL    string `json:"requestUrl"`
		RequestMethod string `json:"requestMethod"`
		RequestBody   struct {
			Parameters struct {
				WaitTime interface{} `json:"waitTime"`
			} `json:"parameters"`
		} `json:"requestBody"`
	} `json:"request"`
	AffectedResources []string `json:"affectedResources"`
}

type Capacities struct {
	Internal map[string]int64 `json:"internal"`
	External map[string]int64 `json:"external"`
	Total    map[string]int64 `json:"total"`
}

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
}

type Port struct {
	PortID             string   `json:"portId"`
	PortType           string   `json:"portType"`
	PortAttributes     []string `json:"portAttributes"`
	PortSpeed          string   `json:"portSpeed"`
	LoopID             string   `json:"loopId"`
	FabricMode         bool     `json:"fabricMode"`
	LunSecuritySetting bool     `json:"lunSecuritySetting"`
}

type PortDetail struct {
	PortID             string   `json:"portId"`
	PortType           string   `json:"portType"`
	PortAttributes     []string `json:"portAttributes"`
	PortSpeed          string   `json:"portSpeed"`
	LoopID             string   `json:"loopId"`
	FabricMode         bool     `json:"fabricMode"`
	PortConnection     string   `json:"portConnection"`
	LunSecuritySetting bool     `json:"lunSecuritySetting"`
	Wwn                string   `json:"wwn"`
	Logins             []struct {
		LoginWwn    string `json:"loginWwn"`
		WwnNickName string `json:"wwnNickName"`
		IsLoggedIn  bool   `json:"isLoggedIn"`
		HostGroupID string `json:"hostGroupId,omitempty"`
	} `json:"logins"`
}

type FCoEPort struct {
	PortID              string   `json:"portId"`
	PortType            string   `json:"portType"`
	PortAttributes      []string `json:"portAttributes"`
	PortSpeed           string   `json:"portSpeed"`
	LoopID              string   `json:"loopId"`
	FabricMode          bool     `json:"fabricMode"`
	PortConnection      string   `json:"portConnection"`
	LunSecuritySetting  bool     `json:"lunSecuritySetting"`
	Wwn                 string   `json:"wwn"`
	StaticMacAddress    string   `json:"staticMacAddress"`
	VLanID              string   `json:"vLanId"`
	DynamicMacAddress   string   `json:"dynamicMacAddress"`
	VirtualPortStatus   string   `json:"virtualPortStatus"`
	VirtualPortID       string   `json:"virtualPortId"`
	FcoeSwitchControlID string   `json:"fcoeSwitchControlId"`
}

type ISCSIPort struct {
	PortID             string   `json:"portId"`
	PortType           string   `json:"portType"`
	PortAttributes     []string `json:"portAttributes"`
	PortSpeed          string   `json:"portSpeed"`
	LoopID             string   `json:"loopId"`
	FabricMode         bool     `json:"fabricMode"`
	LunSecuritySetting bool     `json:"lunSecuritySetting"`
	Logins             []struct {
		LoginIscsiName string `json:"loginIscsiName"`
	} `json:"logins"`
	TCPOption struct {
		Ipv6Mode         bool `json:"ipv6Mode"`
		SelectiveAckMode bool `json:"selectiveAckMode"`
		DelayedAckMode   bool `json:"delayedAckMode"`
		IsnsService      bool `json:"isnsService"`
		TagVLan          bool `json:"tagVLan"`
	} `json:"tcpOption"`
	TCPMtu               int    `json:"tcpMtu"`
	IscsiWindowSize      string `json:"iscsiWindowSize"`
	KeepAliveTimer       int    `json:"keepAliveTimer"`
	TCPPort              string `json:"tcpPort"`
	Ipv4Address          string `json:"ipv4Address"`
	Ipv4Subnetmask       string `json:"ipv4Subnetmask"`
	Ipv4GatewayAddress   string `json:"ipv4GatewayAddress"`
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

type CHAPUser struct {
	ChapUserID      string `json:"chapUserId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostGroupName   string `json:"hostGroupName"`
	ChapUserName    string `json:"chapUserName"`
	WayOfChapUser   string `json:"wayOfChapUser"`
}

type LUN struct {
	LunID           string `json:"lunId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostMode        string `json:"hostMode"`
	Lun             int    `json:"lun"`
	LdevID          int    `json:"ldevId"`
	IsCommandDevice bool   `json:"isCommandDevice"`
	LuHostReserve   struct {
		OpenSystem bool `json:"openSystem"`
		Persistent bool `json:"persistent"`
		PgrKey     bool `json:"pgrKey"`
		Mainframe  bool `json:"mainframe"`
		AcaReserve bool `json:"acaReserve"`
	} `json:"luHostReserve"`
	HostModeOptions []int `json:"hostModeOptions"`
}

type MP struct {
	MpID         int    `json:"mpId"`
	MpLocationID string `json:"mpLocationId"`
	MpUnitID     string `json:"mpUnitId"`
	Ctl          string `json:"ctl"`
}

type CLPR struct {
	ClprID                   int    `json:"clprId"`
	ClprName                 string `json:"clprName"`
	CacheMemoryCapacity      int    `json:"cacheMemoryCapacity"`
	CacheMemoryUsedCapacity  int    `json:"cacheMemoryUsedCapacity"`
	WritePendingDataCapacity int    `json:"writePendingDataCapacity"`
	SideFilesCapacity        int    `json:"sideFilesCapacity"`
	CacheUsageRate           int    `json:"cacheUsageRate"`
	WritePendingDataRate     int    `json:"writePendingDataRate"`
	SideFilesUsageRate       int    `json:"sideFilesUsageRate"`
	ResidentCacheSize        int    `json:"residentCacheSize"`
	NumberOfResidentExtents  int    `json:"numberOfResidentExtents"`
}

type ExternalParityGroup struct {
	ExternalParityGroupID   string `json:"externalParityGroupId"`
	NumOfLdevs              int    `json:"numOfLdevs"`
	UsedCapacityRate        int    `json:"usedCapacityRate"`
	AvailableVolumeCapacity int    `json:"availableVolumeCapacity"`
	EmulationType           string `json:"emulationType"`
	ClprID                  int    `json:"clprId"`
	ExternalProductID       string `json:"externalProductId"`
	Spaces                  []struct {
		PartitionNumber int    `json:"partitionNumber"`
		LdevID          int    `json:"ldevId"`
		Status          string `json:"status"`
		LbaLocation     string `json:"lbaLocation"`
		LbaSize         string `json:"lbaSize"`
	} `json:"spaces"`
}

type ParityGroup struct {
	ParityGroupID                   string `json:"parityGroupId"`
	NumOfLdevs                      int    `json:"numOfLdevs"`
	UsedCapacityRate                int    `json:"usedCapacityRate"`
	AvailableVolumeCapacity         int    `json:"availableVolumeCapacity"`
	RaidLevel                       string `json:"raidLevel"`
	RaidType                        string `json:"raidType"`
	ClprID                          int    `json:"clprId"`
	DriveType                       string `json:"driveType"`
	DriveTypeName                   string `json:"driveTypeName"`
	DriveSpeed                      int    `json:"driveSpeed"`
	TotalCapacity                   int    `json:"totalCapacity"`
	IsAcceleratedCompressionEnabled bool   `json:"isAcceleratedCompressionEnabled"`
}

type Driver struct {
	DriveLocationID string `json:"driveLocationId"`
	DriveTypeName   string `json:"driveTypeName"`
	DriveSpeed      int    `json:"driveSpeed"`
	TotalCapacity   int    `json:"totalCapacity"`
	DriveType       string `json:"driveType"`
	UsageType       string `json:"usageType"`
	Status          string `json:"status"`
	ParityGroupID   string `json:"parityGroupId"`
}

type ResourceGroup struct {
	ResourceGroupID        int      `json:"resourceGroupId"`
	ResourceGroupName      string   `json:"resourceGroupName"`
	LockStatus             string   `json:"lockStatus"`
	LockOwner              string   `json:"lockOwner"`
	LockHost               string   `json:"lockHost"`
	VirtualStorageID       int      `json:"virtualStorageId"`
	LdevIds                []int    `json:"ldevIds"`
	ParityGroupIds         []string `json:"parityGroupIds"`
	ExternalParityGroupIds []string `json:"externalParityGroupIds"`
	PortIds                []string `json:"portIds"`
	HostGroupIds           []string `json:"hostGroupIds"`
}

type UserGroup struct {
	UserGroupObjectID   string   `json:"userGroupObjectId"`
	UserGroupID         string   `json:"userGroupId"`
	RoleNames           []string `json:"roleNames"`
	ResourceGroupIds    []int    `json:"resourceGroupIds"`
	IsBuiltIn           bool     `json:"isBuiltIn"`
	HasAllResourceGroup bool     `json:"hasAllResourceGroup"`
}

type User struct {
	UserObjectID    string   `json:"userObjectId"`
	UserID          string   `json:"userId"`
	Authentication  string   `json:"authentication"`
	UserGroupNames  []string `json:"userGroupNames"`
	IsBuiltIn       bool     `json:"isBuiltIn"`
	IsAccountStatus bool     `json:"isAccountStatus"`
}
