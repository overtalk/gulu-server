package model

type Player struct {
	Id                 int    `json:"id"`
	OpenId             string `json:"open_id"`
	Nickname           string `json:"nickname"`
	Url                string `json:"url"`
	Sex                int    `json:"sex"`
	Strength           int    `json:"strength"`
	Gold               int    `json:"gold"`
	Diamond            int    `json:"diamond"`
	Experience         int    `json:"experience"`
	Treasure           string `json:"treasure"`
	Daily              string `json:"daily"`
	Achievement        string `json:"achievement"`
	Weapon             string `json:"weapon"`
	WeaponEquipped     string `json:"weapon_equipped"`
	Fashion            string `json:"fashion"`
	FashionEquipped    string `json:"fashion_equipped"`
	Olditem            string `json:"olditem"`
	Pve                string `json:"pve"`
	Arena              string `json:"arena"`
	LastPvpRoom        string `json:"last_pvp_room"`
	LastArena          string `json:"last_arena"`
	FreeBox            string `json:"free_box"`
	Dailysign          string `json:"dailysign"`
	CreatedAt          int    `json:"created_at"`
	LastUpdateStrength int    `json:"last_update_strength"`
	Invitation         string `json:"invitation"`
	InvitationCode     int    `json:"invitation_code"`
	ReceivedMail       string `json:"received_mail"`
	IsTest             int    `json:"is_test"`
	FreeLottery        string `json:"free_lottery"`
}

func (*Player) TableName() string {
	return "player"
}
