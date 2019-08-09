package model

type Nodes struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Maxconn  int    `json:"maxconn"`
	IsFull   int    `json:"is_full"` // 0: idle, 1: full
}

func (*Nodes) TableName() string {
	return "nodes"
}
