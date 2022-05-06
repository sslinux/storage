package hcm

import "time"

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

type Session struct {
	SessionID        int       `json:"sessionId"`
	UserID           string    `json:"userId"`
	IPAddress        string    `json:"ipAddress"`
	CreatedTime      time.Time `json:"createdTime"`
	LastAccessedTime time.Time `json:"lastAccessedTime"`
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

type DPPool struct {
	PoolID                                         int    `json:"poolId"`
	PoolStatus                                     string `json:"poolStatus"`
	UsedCapacityRate                               int    `json:"usedCapacityRate"`
	UsedPhysicalCapacityRate                       int    `json:"usedPhysicalCapacityRate"`
	SnapshotCount                                  int    `json:"snapshotCount"`
	PoolName                                       string `json:"poolName"`
	AvailableVolumeCapacity                        int    `json:"availableVolumeCapacity"`
	AvailablePhysicalVolumeCapacity                int    `json:"availablePhysicalVolumeCapacity"`
	TotalPoolCapacity                              int    `json:"totalPoolCapacity"`
	TotalPhysicalCapacity                          int    `json:"totalPhysicalCapacity"`
	NumOfLdevs                                     int    `json:"numOfLdevs"`
	FirstLdevID                                    int    `json:"firstLdevId"`
	WarningThreshold                               int    `json:"warningThreshold"`
	DepletionThreshold                             int    `json:"depletionThreshold"`
	VirtualVolumeCapacityRate                      int    `json:"virtualVolumeCapacityRate"`
	IsMainframe                                    bool   `json:"isMainframe"`
	IsShrinking                                    bool   `json:"isShrinking"`
	LocatedVolumeCount                             int    `json:"locatedVolumeCount"`
	TotalLocatedCapacity                           int    `json:"totalLocatedCapacity"`
	BlockingMode                                   string `json:"blockingMode"`
	TotalReservedCapacity                          int    `json:"totalReservedCapacity"`
	ReservedVolumeCount                            int    `json:"reservedVolumeCount"`
	PoolType                                       string `json:"poolType"`
	DuplicationLdevIds                             []int  `json:"duplicationLdevIds"`
	DuplicationNumber                              int    `json:"duplicationNumber"`
	DataReductionAccelerateCompCapacity            int    `json:"dataReductionAccelerateCompCapacity"`
	DataReductionCapacity                          int    `json:"dataReductionCapacity"`
	DataReductionBeforeCapacity                    int    `json:"dataReductionBeforeCapacity"`
	DataReductionAccelerateCompRate                int    `json:"dataReductionAccelerateCompRate"`
	DuplicationRate                                int    `json:"duplicationRate"`
	CompressionRate                                int    `json:"compressionRate"`
	DataReductionRate                              int    `json:"dataReductionRate"`
	DataReductionAccelerateCompIncludingSystemData struct {
		IsReductionCapacityAvailable bool `json:"isReductionCapacityAvailable"`
		ReductionCapacity            int  `json:"reductionCapacity"`
		IsReductionRateAvailable     bool `json:"isReductionRateAvailable"`
		ReductionRate                int  `json:"reductionRate"`
	} `json:"dataReductionAccelerateCompIncludingSystemData"`
	DataReductionIncludingSystemData struct {
		IsReductionCapacityAvailable bool `json:"isReductionCapacityAvailable"`
		ReductionCapacity            int  `json:"reductionCapacity"`
		IsReductionRateAvailable     bool `json:"isReductionRateAvailable"`
		ReductionRate                int  `json:"reductionRate"`
	} `json:"dataReductionIncludingSystemData"`
	SnapshotUsedCapacity          int  `json:"snapshotUsedCapacity"`
	SuspendSnapshot               bool `json:"suspendSnapshot"`
	CapacitiesExcludingSystemData struct {
		UsedVirtualVolumeCapacity int `json:"usedVirtualVolumeCapacity"`
		CompressedCapacity        int `json:"compressedCapacity"`
		DedupedCapacity           int `json:"dedupedCapacity"`
		ReclaimedCapacity         int `json:"reclaimedCapacity"`
		SystemDataCapacity        int `json:"systemDataCapacity"`
		PreUsedCapacity           int `json:"preUsedCapacity"`
		PreCompressedCapacity     int `json:"preCompressedCapacity"`
		PreDedupredCapacity       int `json:"preDedupredCapacity"`
	} `json:"capacitiesExcludingSystemData"`
}

type ThinImagePool struct {
	PoolID                    int    `json:"poolId"`
	PoolStatus                string `json:"poolStatus"`
	UsedCapacityRate          int    `json:"usedCapacityRate"`
	SnapshotCount             int    `json:"snapshotCount"`
	PoolName                  string `json:"poolName"`
	AvailableVolumeCapacity   int    `json:"availableVolumeCapacity"`
	TotalPoolCapacity         int    `json:"totalPoolCapacity"`
	NumOfLdevs                int    `json:"numOfLdevs"`
	FirstLdevID               int    `json:"firstLdevId"`
	WarningThreshold          int    `json:"warningThreshold"`
	VirtualVolumeCapacityRate int    `json:"virtualVolumeCapacityRate"`
	IsMainframe               bool   `json:"isMainframe"`
	IsShrinking               bool   `json:"isShrinking"`
	PoolType                  string `json:"poolType"`
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
