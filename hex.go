package blockchain

import (
	"bytes"
	"encoding/binary"
)

// ToHex converts number to bytes
func ToHex(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}