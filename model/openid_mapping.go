package model

type Openid_mapping struct {
	PlayerId int    `json:"player_id"`
	OpenId   string `json:"open_id"`
	NodeId   int    `json:"node_id"`
}

func (*Openid_mapping) TableName() string {
	return "openid_mapping"
}
