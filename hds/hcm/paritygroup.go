package hcm

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
