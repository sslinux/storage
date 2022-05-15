package hcm

type LUN struct {
	LunID           string `json:"lunId"`
	PortID          string `json:"portId"`
	HostGroupNumber int    `json:"hostGroupNumber"`
	HostMode        string `json:"hostMode"`
	Lun             int    `json:"lun"`
	LdevID          int    `json:"ldevId"`
	IsCommandDevice bool   `json:"isCommandDevice"`
	NaaID           string `json:"naaId"`
	LuHostReserve   struct {
		OpenSystem bool `json:"openSystem"`
		Persistent bool `json:"persistent"`
		PgrKey     bool `json:"pgrKey"`
		Mainframe  bool `json:"mainframe"`
		AcaReserve bool `json:"acaReserve"`
	} `json:"luHostReserve"`
	HostModeOptions []int `json:"hostModeOptions"`
}
