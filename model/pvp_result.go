package model

type Pvp_result struct {
	PlayerId int    `json:"player_id"`
	Detail   string `json:"detail"`
	IsUsed   int    `json:"is_used"`
	LastTs   int    `json:"last_ts"`
}

func (*Pvp_result) TableName() string {
	return "pvp_result"
}
