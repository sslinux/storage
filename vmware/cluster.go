package vmware

import (
	"fmt"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

type VmwareCluster struct {
	Type              string `json:"type"`
	Value             string `json:"value"`
	Name              string `json:"name"`
	NumHosts          int32  `json:"numHosts"`
	NumEffectiveHosts int32  `json:"numEffectiveHosts"`
	TotalCpu          int32  `json:"totalCpu"`
	TotalMemory       int64  `json:"totalMemory"`
	NumCpuCores       int32  `json:"numCores"`
	NumCpuThreads     int32  `json:"numThreads"`
	EffectiveCpu      int32  `json:"effectiveCpu"`
	EffectiveMemory   int64  `json:"effectiveMemory"`
}

func ClusterInfo(c *vim25.Client) []VmwareCluster {
	var vmClusters []VmwareCluster
	m := view.NewManager(c)
	kind := []string{"ClusterComputeResource"}
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
	if err != nil {
		panic(err)
	}
	defer v.Destroy(ctx)

	var clusters []mo.ClusterComputeResource
	err = v.Retrieve(ctx, kind, []string{"parent", "name", "summary"}, &clusters)
	if err != nil {
		fmt.Println(err)
	}

	for _, cluster := range clusters {
		vmcluster := VmwareCluster{}
		vmcluster.Type = cluster.Self.Type
		vmcluster.Value = cluster.Self.Value
		vmcluster.Name = cluster.Name
		vmcluster.NumHosts = cluster.Summary.GetComputeResourceSummary().NumHosts
		vmcluster.NumEffectiveHosts = cluster.Summary.GetComputeResourceSummary().NumEffectiveHosts
		vmcluster.TotalCpu = cluster.Summary.GetComputeResourceSummary().TotalCpu
		vmcluster.TotalMemory = cluster.Summary.GetComputeResourceSummary().TotalMemory
		vmcluster.NumCpuCores = int32(cluster.Summary.GetComputeResourceSummary().NumCpuCores)
		vmcluster.NumCpuThreads = int32(cluster.Summary.GetComputeResourceSummary().NumCpuThreads)
		vmcluster.EffectiveCpu = cluster.Summary.GetComputeResourceSummary().EffectiveCpu
		vmcluster.EffectiveMemory = cluster.Summary.GetComputeResourceSummary().EffectiveMemory
		vmClusters = append(vmClusters, vmcluster)
	}

	return vmClusters
}
