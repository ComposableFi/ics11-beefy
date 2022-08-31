package types

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/ComposableFi/go-substrate-rpc-client/v4/scale"
)

// Encode scale encodes a data type and returns the scale encoded data as a byte type.
func Encode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := scale.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode decodes `bz` with the scale codec into `target`. `target` should be a pointer.
func Decode(bz []byte, target interface{}) error {
	return scale.NewDecoder(bytes.NewReader(bz)).Decode(target)
}

// DecodeFromHex decodes `str` with the scale codec into `target`. `target` should be a pointer.
func DecodeFromHex(str string, target interface{}) error {
	bz, err := HexDecodeString(str)
	if err != nil {
		return err
	}
	err = scale.NewDecoder(bytes.NewReader(bz)).Decode(target)
	if err != nil {
		return err
	}
	return Decode(bz, target)
}

// HexDecodeString decodes bytes from a hex string. Contrary to hex.DecodeString, this function does not error if "0x"
// is prefixed, and adds an extra 0 if the hex string has an odd length.
func HexDecodeString(s string) ([]byte, error) {
	s = strings.TrimPrefix(s, "0x")

	if len(s)%2 != 0 {
		s = "0" + s
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return b, nil
}