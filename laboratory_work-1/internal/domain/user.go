package domain

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	clientUUID uuid.UUID
	Person
}

type Person struct {
	firstname    string
	lastname     string
	surname      string
	birthdayDate time.Time
}
