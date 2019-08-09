package model

type Player_global_mail struct {
	PlayerId  int `json:"player_id"`
	MailId    int `json:"mail_id"`
	CreatedAt int `json:"created_at"`
}

func (*Player_global_mail) TableName() string {
	return "player_global_mail"
}
