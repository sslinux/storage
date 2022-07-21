package netapp

import (
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"
)

func GetAllAggr() []Aggregate {
	var AllAggregates []Aggregate
	url_postfix := "/api/datacenter/storage/aggregates?order_by=name"

	body := ApiGetRequest(url_postfix)
	aggregates := gjson.Get(string(body), "records").Array()
	for _, aggregate := range aggregates {
		aggr := Aggregate{}
		json.Unmarshal([]byte(aggregate.String()), &aggr)
		AllAggregates = append(AllAggregates, aggr)
	}
	return AllAggregates
}

func (a *Aggregate) CalVolCapacity() {
	AllVols := GetAllVolumes()
	for _, vol := range AllVols {
		if vol.Aggregates[0].Key == a.Key {
			a.CapacityOfVols += vol.Space["size"]
		}
	}
}

func GetAllVolumes() []Volume {
	allVolumes := []Volume{}
	url_postfix := "/api/datacenter/storage/volumes?order_by=name&max_records=10000"
	body := ApiGetRequest(url_postfix)

	volumes := gjson.Get(string(body), "records").Array()
	for _, volume := range volumes {
		vol := Volume{}
		json.Unmarshal([]byte(volume.String()), &vol)
		allVolumes = append(allVolumes, vol)
	}
	// fmt.Println(allVolumes)
	return allVolumes
}

func (v *Volume) GetExportPolicy() []ExportRule {
	url_postfix := v.Nas["export_policy"].Links["self"]["href"]
	v.ExportPolicyName = v.Nas["export_policy"].Name
	body := ApiGetRequest(url_postfix)
	rules := gjson.Get(string(body), "rules").Array()
	ExportRules := []ExportRule{}
	for _, rule := range rules {
		var export_rule = ExportRule{}
		json.Unmarshal([]byte(rule.String()), &export_rule)
		export_rule.Merge()
		export_rule.VolName = v.Name
		export_rule.ExportPolicyName = v.ExportPolicyName
		ExportRules = append(ExportRules, export_rule)
	}

	return ExportRules
}

func (e *ExportRule) Merge() {
	var ips []string
	for _, ip := range e.Clients {
		ips = append(ips, ip["match"])
	}
	e.StrClients = strings.Join(ips, ",")
	e.StrProtocols = strings.Join(e.Protocols, ",")
	e.StrSuperuser = strings.Join(e.Superuser, ",")
	e.StrRORule = strings.Join(e.RO_Rule, ",")
	e.StrRWRule = strings.Join(e.RW_Rule, ",")
}
