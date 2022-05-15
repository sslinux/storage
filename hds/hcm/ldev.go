package hcm

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

