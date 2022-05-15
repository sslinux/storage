package hcm

import "time"

type Job struct {
	JobID         int       `json:"jobId"`
	Self          string    `json:"self"`
	UserID        string    `json:"userId"`
	Status        string    `json:"status"`
	State         string    `json:"state"`
	CreatedTime   time.Time `json:"createdTime"`
	UpdatedTime   time.Time `json:"updatedTime"`
	CompletedTime time.Time `json:"completedTime"`
	Request       struct {
		RequestURL    string `json:"requestUrl"`
		RequestMethod string `json:"requestMethod"`
		RequestBody   struct {
			Parameters struct {
				WaitTime interface{} `json:"waitTime"`
			} `json:"parameters"`
		} `json:"requestBody"`
	} `json:"request"`
	AffectedResources []string `json:"affectedResources"`
}
