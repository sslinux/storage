package hcm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

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

func GetDeviceIDBySN(sn int64) StorageSystem {
	var targetStorage StorageSystem
	storages := GetAllStorages()
	for _, storage := range storages {
		if storage.SerialNumber == sn {
			targetStorage = storage
		}
	}
	return targetStorage
}

type StorageSystem struct {
	StorageDeviceID string `json:"storageDeviceId"`
	Model           string `json:"model"`
	SerialNumber    int64  `json:"serialNumber"`
	SvpIP           string `json:"svpIp"`
}

type Capacities struct {
	Internal map[string]int64 `json:"internal"`
	External map[string]int64 `json:"external"`
	Total    map[string]int64 `json:"total"`
}

type CHAPUser struct {
	ChapUserID      string `json:"chapUserId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostGroupName   string `json:"hostGroupName"`
	ChapUserName    string `json:"chapUserName"`
	WayOfChapUser   string `json:"wayOfChapUser"`
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