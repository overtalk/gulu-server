package serializer

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigDefault

func jsonMarshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return addSerializationProtoID(data, jsonSerializer), nil
}

func jsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
