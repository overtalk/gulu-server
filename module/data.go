package module

import (
	"encoding/json"
)

type PostResponse struct {
	ErrCode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

func (p PostResponse) Encode() string {
	d, err := json.Marshal(p)
	if err != nil {
		p.ErrCode = 1001
		p.Msg = "反序列化错误"
	}
	return string(d)
}

// PvpInfo
type PvpInfo struct {
	ID    int `json:"id"`
	Arena int `json:"arena"`
	Cup   int `json:"cup"`
}

func PvpInfoReq(data []byte) *PvpInfo {
	p := &PvpInfo{}
	if err := json.Unmarshal(data, &p); err != nil {
		return &PvpInfo{}
	}
	return p
}

// basic info
type BasicInfo struct {
	ID         int `json:"id"`
	Gold       int `json:"gold"`
	Diamond    int `json:"diamond"`
	Experience int `json:"experience"`
	Strength   int `json:"strength"`
}

func BasicInfoReq(data []byte) *BasicInfo {
	p := &BasicInfo{}
	if err := json.Unmarshal(data, &p); err != nil {
		return &BasicInfo{}
	}
	return p
}

// common operation
type CommonOp struct {
	OP int `json:"op"`
}

func CommonOpReq(data []byte) *CommonOp {
	p := &CommonOp{}
	if err := json.Unmarshal(data, &p); err != nil {
		return &CommonOp{}
	}
	return p
}
