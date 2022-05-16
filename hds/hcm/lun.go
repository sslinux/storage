package hcm

type LUN struct {
	LunID           string `json:"lunId"`
	PortID          string `json:"portId"`
	HostGroupNumber int64  `json:"hostGroupNumber"`
	HostMode        string `json:"hostMode"`
	Lun             int64  `json:"lun"`
	LdevID          int64  `json:"ldevId"`
	IsCommandDevice bool   `json:"isCommandDevice"`
	NaaID           string `json:"naaId"`
	LuHostReserve   struct {
		OpenSystem bool `json:"openSystem"`
		Persistent bool `json:"persistent"`
		PgrKey     bool `json:"pgrKey"`
		Mainframe  bool `json:"mainframe"`
		AcaReserve bool `json:"acaReserve"`
	} `json:"luHostReserve"`
	HostModeOptions []int64 `json:"hostModeOptions"`
}

func (lun *LUN) GetLunNaaID(session Session) (string, error) {
	tmpLdev := GetSpecifyLDEV(session, lun.LdevID)
	lun.NaaID = tmpLdev.NaaId
	return tmpLdev.NaaId, nil
}
