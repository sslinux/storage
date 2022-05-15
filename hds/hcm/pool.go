package hcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

type Pool struct {
	gorm.Model
	SN                        int64
	PoolID                    int64  `json:"poolId"`
	PoolStatus                string `json:"poolStatus"`
	UsedCapacityRate          int64  `json:"usedCapacityRate"`
	SnapshotCount             int64  `json:"snapshotCount"`
	PoolName                  string `json:"poolName"`
	AvailableVolumeCapacity   int64  `json:"availableVolumeCapacity"`
	TotalPoolCapacity         int64  `json:"totalPoolCapacity"`
	NumOfLdevs                int64  `json:"numOfLdevs"`
	FirstLdevID               int64  `json:"firstLdevId"`
	WarningThreshold          int64  `json:"warningThreshold"`
	VirtualVolumeCapacityRate int64  `json:"virtualVolumeCapacityRate"`
	IsMainframe               bool   `json:"isMainframe"`
	IsShrinking               bool   `json:"isShrinking"`
	PoolType                  string `json:"poolType"`
	SaveEffect                float64
}

type DPPool struct {
	Pool
	UsedPhysicalCapacityRate        int64  `json:"usedPhysicalCapacityRate"`
	AvailablePhysicalVolumeCapacity int64  `json:"availablePhysicalVolumeCapacity"`
	TotalPhysicalCapacity           int64  `json:"totalPhysicalCapacity"`
	UsedPhysicalCapacity            int64  `json:"usedPhysicalCapacity"`
	DepletionThreshold              int64  `json:"depletionThreshold"`
	LocatedVolumeCount              int64  `json:"locatedVolumeCount"`
	TotalLocatedCapacity            int64  `json:"totalLocatedCapacity"`
	BlockingMode                    string `json:"blockingMode"`
	TotalReservedCapacity           int64  `json:"totalReservedCapacity"`
	ReservedVolumeCount             int64  `json:"reservedVolumeCount"`
	// DuplicationLdevIds                  []int64 `json:"duplicationLdevIds"`
	DuplicationNumber                   int64 `json:"duplicationNumber"`
	DataReductionAccelerateCompCapacity int64 `json:"dataReductionAccelerateCompCapacity"`
	DataReductionCapacity               int64 `json:"dataReductionCapacity"`
	DataReductionBeforeCapacity         int64 `json:"dataReductionBeforeCapacity"`
	DataReductionAccelerateCompRate     int64 `json:"dataReductionAccelerateCompRate"`
	DuplicationRate                     int64 `json:"duplicationRate"`
	CompressionRate                     int64 `json:"compressionRate"`
	DataReductionRate                   int64 `json:"dataReductionRate"`
	// DataReductionAccelerateCompIncludingSystemData struct {
	// 	IsReductionCapacityAvailable bool  `json:"isReductionCapacityAvailable"`
	// 	ReductionCapacity            int64 `json:"reductionCapacity"`
	// 	IsReductionRateAvailable     bool  `json:"isReductionRateAvailable"`
	// 	ReductionRate                int64 `json:"reductionRate"`
	// } `json:"dataReductionAccelerateCompIncludingSystemData"`
	// DataReductionIncludingSystemData struct {
	// 	IsReductionCapacityAvailable bool  `json:"isReductionCapacityAvailable"`
	// 	ReductionCapacity            int64 `json:"reductionCapacity"`
	// 	IsReductionRateAvailable     bool  `json:"isReductionRateAvailable"`
	// 	ReductionRate                int64 `json:"reductionRate"`
	// } `json:"dataReductionIncludingSystemData"`
	// SnapshotUsedCapacity          int64 `json:"snapshotUsedCapacity"`
	// SuspendSnapshot               bool  `json:"suspendSnapshot"`
	// CapacitiesExcludingSystemData struct {
	// 	UsedVirtualVolumeCapacity int64 `json:"usedVirtualVolumeCapacity"`
	// 	CompressedCapacity        int64 `json:"compressedCapacity"`
	// 	DedupedCapacity           int64 `json:"dedupedCapacity"`
	// 	ReclaimedCapacity         int64 `json:"reclaimedCapacity"`
	// 	SystemDataCapacity        int64 `json:"systemDataCapacity"`
	// 	PreUsedCapacity           int64 `json:"preUsedCapacity"`
	// 	PreCompressedCapacity     int64 `json:"preCompressedCapacity"`
	// 	PreDedupredCapacity       int64 `json:"preDedupredCapacity"`
	// } `json:"capacitiesExcludingSystemData"`
}

type ThinImagePool struct {
	Pool
}

func GetPools(session *Session, poolType, detailInfoType string) interface{} {
	var Parameters = map[string]string{}

	if poolType != "" {
		Parameters["poolType"] = poolType
	}

	if detailInfoType != "" {
		Parameters["detailInfoType"] = detailInfoType
	}

	resp, err := session.Request("GET", "/pools", Parameters, nil, nil)
	if err != nil {
		log.Printf("GetPools Error:%v\n", err)
		fmt.Println(err)
	}
	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read response body error:%s\n", err)
	}

	switch poolType {
	case "DP":
		var pools []DPPool
		for _, p := range gjson.Get(string(byteBody), "data").Array() {
			pool := DPPool{}
			json.Unmarshal([]byte(p.String()), &pool)
			pools = append(pools, pool)
		}
		return pools
	case "HTI":
		var pools []ThinImagePool
		for _, p := range gjson.Get(string(byteBody), "data").Array() {
			pool := ThinImagePool{}
			json.Unmarshal([]byte(p.String()), &pool)
			pools = append(pools, pool)
		}
		return pools
	case "":
		var pools []Pool
		for _, p := range gjson.Get(string(byteBody), "data").Array() {
			pool := Pool{}
			json.Unmarshal([]byte(p.String()), &pool)
			pools = append(pools, pool)
		}
		return pools
	default:
		log.Printf("poolType error:%s\n", poolType)
	}
	return nil
}
