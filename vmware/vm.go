package vmware

import (
	"fmt"
	"strings"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

type TargetVM struct {
	Group         string   `json:"group"`
	VMID          string   `json:"vmid"`
	VMName        string   `json:"name"`
	VMIP          string   `json:"ip"`
	DatastoreName []string `json:"datastoreName"`
	DataStoreID   []string `json:"datastoreID"`
	HostID        string   `json:"hostId"`
	HostIP        string   `json:"hostIP"`
}

func GetAllVM(c *vim25.Client) {
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
		for _, d := range vm.Config.DatastoreUrl {
			tvm.DatastoreName = append(tvm.DatastoreName, d.Name)
		}

		tvm.DataStoreID = []string{"待补充"}
		tvm.HostIP = "待增加取物理主机IP的方法"

		fmt.Println(tvm)
	}

	// bytevm, _ := json.Marshal(vms[0])
	// fmt.Println(string(bytevm))

	// fmt.Println(vms)
}
