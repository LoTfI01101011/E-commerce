package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time
}
