package model

type Single_mail struct {
	Id        int    `json:"id"`
	PlayerId  int    `json:"player_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Icon      int    `json:"icon"`
	Award     string `json:"award"`
	CreatedAt int    `json:"created_at"`
	Indate    int    `json:"indate"`
	Image     string `json:"image"`
}

func (*Single_mail) TableName() string {
	return "single_mail"
}
