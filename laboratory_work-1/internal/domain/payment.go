package domain

import "github.com/google/uuid"

type Transaction struct {
	id              int
	srcAccountUUID  uuid.UUID
	destAccountUUID uuid.UUID
	status          PaymentStatus
	value           Money
	txType          TransactionType
}

type PaymentStatus int

const (
	Initiated PaymentStatus = iota
	Pending
	Declined
	Failed
	Completed
)

type TransactionType int

const (
	Transfer TransactionType = iota + 1
	Deposit
	Withdraw
)

func NewTransaction(id int, srcAccountUUID uuid.UUID, destAccountUUID uuid.UUID, value Money, txType TransactionType) Transaction {
	return Transaction{
		id:              id,
		srcAccountUUID:  srcAccountUUID,
		destAccountUUID: destAccountUUID,
		status:          Initiated,
		value:           value,
		txType:          txType,
	}
}

func (t Transaction) ID() int {
	return t.id
}

func (t Transaction) SourceAccountUUID() uuid.UUID {
	return t.srcAccountUUID
}

func (t Transaction) DestinationAccountUUID() uuid.UUID {
	return t.destAccountUUID
}

func (t Transaction) Status() PaymentStatus {
	return t.status
}

func (t Transaction) Value() Money {
	return t.value
}

func (t Transaction) Type() TransactionType {
	return t.txType
}

func (t *Transaction) SetStatus(status PaymentStatus) {
	t.status = status
}
