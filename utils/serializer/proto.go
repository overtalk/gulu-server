package serializer

import (
	"encoding/binary"
)

type serializerProtoID uint16

const protoLen = 2

const (
	invalidSerializer serializerProtoID = iota
	jsonSerializer
)

func addSerializationProtoID(data []byte, protoID serializerProtoID) []byte {
	buf := make([]byte, len(data)+protoLen)
	binary.BigEndian.PutUint16(buf[:protoLen], uint16(protoID))
	copy(buf[protoLen:], data)
	return buf
}

func getSerializerProtoID(data []byte) (serializerProtoID, error) {
	if len(data) < protoLen {
		return invalidSerializer, ErrInValidSerializer
	}
	return serializerProtoID(binary.BigEndian.Uint16(data[:protoLen])), nil
}
