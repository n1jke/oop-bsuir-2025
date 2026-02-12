package domain

import "github.com/google/uuid"

type Transaction struct {
	id              int
	srcAccountUUID  uuid.UUID
	destAccountUUID uuid.UUID
	status          PaymentStatus
	value           Money
}

type PaymentStatus int

const (
	Initiated PaymentStatus = iota
	Pending
	Declined
	Failed
)
