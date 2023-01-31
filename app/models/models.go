package models

import (
	"time"

	"github.com/bytedance/sonic"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"                         json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:timezone('utc'::text, CURRENT_TIMESTAMP);index" json:"updatedAt"`
}

func Update(modelPtr any, schema any) error {
	schemaBytes, err := sonic.Marshal(schema)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(schemaBytes, modelPtr)
}
