package netapp

type Aggregate struct {
	Key            string                    `json:"key"`
	Uuid           string                    `json:"uuid"`
	Name           string                    `json:"name"`
	CLuster        Cluster                   `json:"cluster"`
	NOde           Node                      `json:"node"`
	Space          map[string]map[string]int `json:"space"`
	CapacityOfVols int64
}

type Cluster struct {
	Key  string `json:"key"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Node struct {
	Key  string `json:"key"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

type Svm struct {
	Key   string                       `json:"key"`
	Uuid  string                       `json:"uuid"`
	Name  string                       `json:"name"`
	Links map[string]map[string]string `json:"_links"`
}

type Volume struct {
	Key              string                  `json:"key"`
	Uuid             string                  `json:"uuid"`
	Name             string                  `json:"name"`
	Cluster          Cluster                 `json:"cluster"`
	Aggregates       []Aggregate             `json:"aggregates"`
	Space            map[string]int64        `json:"space"`
	Snapmirror       map[string]bool         `json:"snapmirror"`
	Svm              Svm                     `json:"svm"`
	Nas              map[string]ExportPolicy `json:"nas"`
	ExportPolicyName string
}

type ExportPolicy struct {
	Key   string                       `json:"key"`
	Id    string                       `json:"id"`
	Name  string                       `json:"name"`
	Links map[string]map[string]string `json:"_links"`
}

type ExportRule struct {
	Clients          []map[string]string `json:"clients"`
	Index            int64               `json:"index"`
	Protocols        []string            `json:"protocols"`
	Superuser        []string            `json:"superuser"`
	Anonymous_User   string              `json:"anonymous_user"`
	RO_Rule          []string            `json:"ro_rule"`
	RW_Rule          []string            `json:"rw_rule"`
	StrClients       string
	StrProtocols     string
	StrSuperuser     string
	StrRORule        string
	StrRWRule        string
	VolName          string
	ExportPolicyName string
}
