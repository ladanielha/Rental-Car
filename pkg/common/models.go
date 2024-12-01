package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelsWithID struct {
	ID        string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *ModelsWithID) generateID(prefix string) string {
	return prefix + "_" + uuid.New().String()
}

func (m *ModelsWithID) GenerateUUID(prefix string) {
	m.ID = m.generateID(prefix)
}
