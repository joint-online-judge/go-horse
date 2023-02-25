package model

import (
	"github.com/joint-online-judge/go-horse/pkg/json"
)

func Update(modelPtr any, schema any) error {
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	return json.Unmarshal(schemaBytes, modelPtr)
}
