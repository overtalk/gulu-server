package protocol

// PvpInfo
type PvpInfo struct {
	ID    int `json:"id"`
	Arena int `json:"arena"`
	Cup   int `json:"cup"`
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
