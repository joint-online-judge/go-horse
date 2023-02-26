package json

import (
	"github.com/bytedance/sonic"
)

func Marshal(val any) ([]byte, error) {
	return sonic.Marshal(val)
}

func Unmarshal(buf []byte, val any) error {
	return sonic.Unmarshal(buf, val)
}
