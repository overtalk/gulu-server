package protocol

// LoginReq is login request struct
type LoginReq struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// login response from client
type LoginResponse struct {
	Response
	Status string `json:"status"`
	Type   string `json:"type"`
	Token  string `json:"token"`
}

type UserInfo struct {
	Name             string `json:"name"`
	Avatar           string `json:"avatar"`
	CurrentAuthority string `json:"currentAuthority"`
}

type CurrentResponse struct {
	Response
	UserInfo UserInfo `json:"userInfo"`
}
