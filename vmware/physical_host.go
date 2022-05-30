package vmware

import (
	"context"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
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
