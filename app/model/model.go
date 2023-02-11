package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/joint-online-judge/go-horse/pkg/json"
)

type Base struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
}

func Update(modelPtr any, schema any) error {
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	return json.Unmarshal(schemaBytes, modelPtr)
}
