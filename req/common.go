package req

import (
	"encoding/json"
)

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
