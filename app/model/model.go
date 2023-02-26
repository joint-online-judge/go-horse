package model

import (
	"github.com/joint-online-judge/go-horse/pkg/json"
)

func Update[Model any, Schema any](model *Model, schema Schema) error {
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	return json.Unmarshal(schemaBytes, model)
}
