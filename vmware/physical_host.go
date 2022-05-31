package vmware

import (
	"context"
	"log"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
)

var U = &url.URL{}
var ctx = context.Background()

func init() {

	// for cloud1
	// var user = "nbuprd@vsphere.local"
	// var password = "CrcloudNbu@-2020"
	// var ip = "192.168.231.241"

	// for chuantong prod
	// var user = `crhd0a\xiongguiyin`
	// var password = "6524198As!@"
	// var ip = "10.0.66.11"

	// 传统资源池测试,版本太低，报错；
	// var user = `crhd0a\xiongguiyin`
	// var password = "6524198As!@"
	// var ip = "10.0.66.19"

	// 沙河灾备
	var user = `crhd0a\xiongguiyin`
	var password = "6524198As!@"
	var ip = "10.0.68.171"

	U.User = url.UserPassword(user, password)
	U.Host = ip
	U.Scheme = "https"
	U.Path = "sdk"
}

func NewClient() *vim25.Client {
	// fmt.Println(U)
	client, err := govmomi.NewClient(ctx, U, true)
	if err != nil {
		// panic(err)
		log.Printf("NewClient error:%s\n", err)
	}
	return client.Client
}
