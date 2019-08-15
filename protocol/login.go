package protocol

// LoginReq is login request struct
type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Type     string `json:"account"`
}
