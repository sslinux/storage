package vmware

import (
	"context"
	"log"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
)

var ctx = context.Background()

func NewClient(ctx context.Context, U *url.URL) *vim25.Client {
	// fmt.Println(U)
	client, err := govmomi.NewClient(ctx, U, true)
	if err != nil {
		// panic(err)
		log.Printf("NewClient error:%s\n", err)
	}
	return client.Client
}
