package protocol

type Player struct {
	ID         int   `json:"id"`
	Gold       int   `json:"gold"`
	Diamond    int   `json:"diamond"`
	Experience int   `json:"experience"`
	Strength   int   `json:"strength"`
	Arena      int   `json:"arena"`
	Score      []int `json:"score"`
}

type PlayerQuery struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

// PvpInfo
type PvpInfo struct {
	ID    int   `json:"id"`
	Arena int   `json:"arena"`
	Socre []int `json:"score"`
}

// basic info
type BasicInfo struct {
	ID         int `json:"id"`
	Gold       int `json:"gold"`
	Diamond    int `json:"diamond"`
	Experience int `json:"experience"`
	Strength   int `json:"strength"`
}

// common operation
type CommonOP struct {
	ID int `json:"id"`
	OP int `json:"op"`
}

// post response from client
type PostResponse struct {
	Response
	Player Player `json:"player"`
}
