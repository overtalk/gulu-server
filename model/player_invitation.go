package model

type Player_invitation struct {
	PlayerId  int `json:"player_id"`
	FriendId  int `json:"friend_id"`
	CreatedAt int `json:"created_at"`
}

func (*Player_invitation) TableName() string {
	return "player_invitation"
}
