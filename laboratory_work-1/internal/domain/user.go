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

func NewClient(clientUUID uuid.UUID, person Person) Client {
	return Client{
		clientUUID: clientUUID,
		Person:     person,
	}
}

func (c *Client) ClientUUID() uuid.UUID {
	return c.clientUUID
}
