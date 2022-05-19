package vmware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

var U = &url.URL{}
var ctx = context.Background()

func init() {
	var user = "nbuprd@vsphere.local"
	var password = "CrcloudNbu@-2020"
	var ip = "192.168.231.241"
	U.User = url.UserPassword(user, password)
	U.Host = ip
	U.Scheme = "https"
	U.Path = "sdk"
}

func NewClient() *vim25.Client {
	client, err := govmomi.NewClient(ctx, U, true)
	if err != nil {
		panic(err)
	}
	return client.Client
}

func FindDatastore(c *vim25.Client) {
	m := view.NewManager(c)
	kind := []string{"Datastore"}
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(ctx)

	var stores []mo.Datastore
	err = v.Retrieve(ctx, kind, []string{"name", "parent", "host", "vm", "DatastoreHostMount"}, &stores)
	if err != nil {
		fmt.Println(err)
	}
	for _, store := range stores {
		byteStore, _ := json.Marshal(store)
		fmt.Println(string(byteStore))
	}

}
