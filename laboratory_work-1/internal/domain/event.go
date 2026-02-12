package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Name        string
	AccountUUID uuid.UUID
	OccurredAt  time.Time
}
