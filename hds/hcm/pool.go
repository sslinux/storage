package hcm

import (
	"fmt"
	"log"
)

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

func GetPools(session *Session, poolType, detailInfoType string) {
	var Parameters = map[string]string{}

	if poolType != "" {
		Parameters["poolType"] = poolType
	}

	if detailInfoType != "" {
		Parameters["detailInfoType"] = detailInfoType
	}

	err := session.Request("GET", "/pools", Parameters, nil, nil)
	if err != nil {
		log.Printf("GetPools Error:%v\n", err)
		fmt.Println(err)
	}
}
