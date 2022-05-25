package vmware

type TargetHost struct {
	HostID         string
	Name           string
	WWNN           []string
	WWPN           []string
	ScsiLuns       []ScsiLUN
	NumOfVM        int64    `json:"numOfVM"`
	NumOfDatastore int64    `json:"numOfDatastore"`
	VMs            []string `json:"vms"`
	Datastores     []string `json:"datastores"`
}

type ScsiLUN struct {
	CanonicalName string `json:"canonicalName"`
	Vender        string `json:"vender"`
	Model         string `json:"model"`
	LocalDisk     bool   `json:"localDisk"`
}
