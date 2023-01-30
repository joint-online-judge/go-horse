package models

import (
	"github.com/bytedance/sonic"
)

func Update(modelPtr any, schema any) error {
	schemaBytes, err := sonic.Marshal(schema)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(schemaBytes, modelPtr)
}
