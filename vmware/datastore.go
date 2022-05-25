package vmware

type TargetDatastore struct {
	Datastore_ID string
	Name         string `json:"Name"`
	FreeSpace    int64  `json:"FreeSpace"`
	Capacity     int64  `json:"Capacity"`
	BlockSize    int64  `json:"BlockSize"`
	Uuid         string `json:"Uuid"`
	DiskName     string `json:"DiskName"`
	Hosts        []string
	Vms          []string
}
