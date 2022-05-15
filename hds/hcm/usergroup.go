package hcm

type UserGroup struct {
	UserGroupObjectID   string   `json:"userGroupObjectId"`
	UserGroupID         string   `json:"userGroupId"`
	RoleNames           []string `json:"roleNames"`
	ResourceGroupIds    []int    `json:"resourceGroupIds"`
	IsBuiltIn           bool     `json:"isBuiltIn"`
	HasAllResourceGroup bool     `json:"hasAllResourceGroup"`
}

type User struct {
	UserObjectID    string   `json:"userObjectId"`
	UserID          string   `json:"userId"`
	Authentication  string   `json:"authentication"`
	UserGroupNames  []string `json:"userGroupNames"`
	IsBuiltIn       bool     `json:"isBuiltIn"`
	IsAccountStatus bool     `json:"isAccountStatus"`
}


