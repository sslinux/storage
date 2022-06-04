package vmware

import (
	"fmt"
	"strings"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"gorm.io/gorm"
)

type TargetVM struct {
	gorm.Model
	Group         string `json:"group"`
	VMID          string `json:"vmid"`
	VMName        string `json:"name"`
	VMIP          string `json:"ip"`
	DatastoreName string `json:"datastoreName"`
	DataStoreID   string `json:"datastoreID"`
	HostID        string `json:"hostId"`
	HostIP        string `json:"hostIP"`
}

var AllHostIDMapName map[string]string
var AllDataStoreIDMapName map[string]string

func GetAllVM(c *vim25.Client) []TargetVM {
	AllTargetHosts := GetAllHost(c)
	AllHostIDMapName = GetHostIDMapName(AllTargetHosts)

	AllTargetDatastores := GetAllDatastore(c)
	AllDataStoreIDMapName = GetAllDatastoreIDMapName(AllTargetDatastores)

	var tvms []TargetVM
	m := view.NewManager(c)
	kind := []string{"VirtualMachine"}
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(ctx)

	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, kind, []string{"parent", "name", "summary", "runtime", "storage", "config", "datastore"}, &vms)
	if err != nil {
		fmt.Println(err)
	}

	for _, vm := range vms {
		tvm := TargetVM{}
		tvm.VMName = vm.Name
		if ok := strings.HasPrefix(vm.Name, "Volume-") || strings.HasPrefix(vm.Name, "volume-"); ok {
			continue
		}
		tvm.VMID = vm.Self.Value
		tvm.Group = vm.Parent.Value
		tvm.HostID = vm.Summary.Runtime.Host.Value
		tvm.VMIP = vm.Summary.Guest.IpAddress
		tvm.HostIP = AllHostIDMapName[tvm.HostID]
		for _, d := range vm.Config.DatastoreUrl {
			tvm.DatastoreName = d.Name
			tvm.DataStoreID = AllDataStoreIDMapName[d.Name]
			tvms = append(tvms, tvm)
		}
	}

	// bytevm, _ := json.Marshal(vms[0])
	// fmt.Println(string(bytevm))

	// fmt.Println(vms)
	return tvms
}
