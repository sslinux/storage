package vmware

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

// var ctx = context.Background()

type TargetHost struct {
	HostID         string
	Name           string
	WWNN           []string
	WWPN           []string
	NumOfScsiLun   int
	ScsiLuns       []ScsiLUN
	NumOfVM        int      `json:"numOfVM"`
	NumOfDatastore int      `json:"numOfDatastore"`
	VMs            []string `json:"vms"`
	Datastores     []string `json:"datastores"`
}

type ScsiLUN struct {
	CanonicalName string `json:"canonicalName"`
	Vender        string `json:"vender"`
	Model         string `json:"model"`
	LocalDisk     bool   `json:"localDisk"`
	UUID          string `json:"uuid"`
	NumOfPath     int    `json:"numOfPath"`
}

func GetAllHost(c *vim25.Client) []TargetHost {
	m := view.NewManager(c)
	kind := []string{"HostSystem"}
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(ctx)

	var hosts []mo.HostSystem
	err = v.Retrieve(ctx, kind, []string{"parent", "name", "config", "datastore", "network", "vm"}, &hosts)
	if err != nil {
		fmt.Println(err)
	}

	var thosts []TargetHost
	for _, host := range hosts {
		thost, err := GetHostInfo(&host)
		if err != nil {
			log.Printf("GetHostInfo error: %v", err)
		}
		thosts = append(thosts, thost)
		// fmt.Println(thost)
	}
	return thosts
}

func GetHostIDMapName(thosts []TargetHost) map[string]string {
	hostIDMapName := make(map[string]string)
	for _, thost := range thosts {
		hostIDMapName[thost.HostID] = thost.Name
	}
	return hostIDMapName
}

func GetHostInfo(host *mo.HostSystem) (TargetHost, error) {
	thost := TargetHost{}
	thost.HostID = host.Self.Value
	thost.Name = host.Name

	// 获取主机的WWNN和WWPN
	adapters := host.Config.StorageDevice.HostBusAdapter
	// printStr, _ := json.Marshal(adapters)
	// fmt.Println(string(printStr))
	for _, adapter := range adapters {
		strAdapter, _ := json.Marshal(adapter)
		// ok := strings.Contains(gjson.Get(string(strAdapter), "Model").String(), "Fibre Channel Adapter")
		// if ok {
		if gjson.Get(string(strAdapter), "PortWorldWideName").Int() != 0 {
			thost.WWPN = append(thost.WWPN, strconv.FormatInt(gjson.Get(string(strAdapter), "PortWorldWideName").Int(), 16))
			fmt.Println(thost.Name, strconv.FormatInt(gjson.Get(string(strAdapter), "PortWorldWideName").Int(), 16))
			thost.WWNN = append(thost.WWNN, strconv.FormatInt(gjson.Get(string(strAdapter), "NodeWorldWideName").Int(), 16))
		}
		// }
	}

	// 获取主机的ScsiLUNs
	thost.NumOfScsiLun = len(host.Config.StorageDevice.ScsiLun)
	for _, lun := range host.Config.StorageDevice.ScsiLun {
		tlun := ScsiLUN{}
		tlun.CanonicalName = lun.GetScsiLun().CanonicalName
		tlun.Vender = lun.GetScsiLun().Vendor
		tlun.Model = lun.GetScsiLun().Model
		tlun.UUID = lun.GetScsiLun().Uuid

		for _, LunPath := range host.Config.StorageDevice.MultipathInfo.Lun {
			if tlun.UUID != LunPath.Id {
				continue
			}
			tlun.NumOfPath = len(LunPath.Path)
		}
		thost.ScsiLuns = append(thost.ScsiLuns, tlun)
	}

	// 获取虚拟机的信息
	thost.NumOfVM = len(host.Vm)
	for _, vm := range host.Vm {
		thost.VMs = append(thost.VMs, vm.Value)
	}

	// 获取Datastore的信息
	thost.NumOfDatastore = len(host.Datastore)
	for _, datastore := range host.Datastore {
		thost.Datastores = append(thost.Datastores, datastore.Value)
	}

	return thost, nil
}
