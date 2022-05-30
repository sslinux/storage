package vmware

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

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

func GetAllDatastore(c *vim25.Client) []TargetDatastore {
	m := view.NewManager(c)
	kind := []string{"Datastore"}
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(ctx)

	var stores []mo.Datastore
	err = v.Retrieve(ctx, kind, []string{"summary", "name", "info", "host", "vm"}, &stores)
	if err != nil {
		fmt.Println(err)
	}

	var allDatastores []TargetDatastore
	for _, store := range stores {
		targetDatastore := TargetDatastore{}
		targetDatastore.Datastore_ID = store.Self.Value
		targetDatastore.Name = store.Name
		targetDatastore.FreeSpace = store.Info.GetDatastoreInfo().FreeSpace
		byteInfo, _ := json.Marshal(store.Info)

		targetDatastore.Capacity = gjson.Get(string(byteInfo), "Vmfs.Capacity").Int()
		targetDatastore.BlockSize = gjson.Get(string(byteInfo), "Vmfs.BlockSize").Int()
		targetDatastore.Uuid = gjson.Get(string(byteInfo), "Vmfs.Uuid").String()

		if len(gjson.Get(string(byteInfo), "Vmfs.Extent").Array()) == 1 {
			tmpStr := gjson.Get(string(byteInfo), "Vmfs.Extent").Array()[0].String()
			targetDatastore.DiskName = gjson.Get(tmpStr, "DiskName").String()
		}

		// fmt.Println(string(byteInfo))
		// NaaID := gjson.Get(string(byteInfo), "Vmfs.Extent").String()
		// fmt.Println(store.Self.Value, store.Name, NaaID)
		// byteprint, _ := json.Marshal(store.Info)
		// fmt.Println(string(byteprint))

		for _, h := range store.Host {
			targetDatastore.Hosts = append(targetDatastore.Hosts, h.Key.Value)
			// fmt.Printf("%s\t", h.Key.Value)
		}
		// fmt.Println()
		for _, v := range store.Vm {
			targetDatastore.Vms = append(targetDatastore.Vms, v.Value)
			// fmt.Printf("%s\t", v.Value)
		}
		// byteStore, _ := json.Marshal(store)
		// fmt.Println(string(byteStore))
		// fmt.Println()
		allDatastores = append(allDatastores, targetDatastore)
	}
	return allDatastores
}
