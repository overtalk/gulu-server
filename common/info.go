package common

import (
	"encoding/json"
)

// basic info
type BasicInfo struct {
	ID         int `json:"id"`
	Gold       int `json:"gold"`
	Diamond    int `json:"diamond"`
	Experience int `json:"experience"`
	Strength   int `json:"strength"`
}

func TurnBasicInfo(data []byte) *BasicInfo {
	p := &BasicInfo{}
	if err := json.Unmarshal(data, &p); err != nil {
		return &BasicInfo{}
	}
	return p
}
