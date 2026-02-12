package application

import (
	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

type PaymentService interface {
	Deposit(m domain.Money, aUUID uuid.UUID) error
	Withdraw(m domain.Money, aUUID uuid.UUID) error
	ProcessTransaction(t domain.Transaction) error
}

type AccountRepository interface {
	Create(a PaymentAccount) error
	ChangeStatus(aUUID uuid.UUID, s domain.AccountStatus) error
}

type EventService interface {
	Publish(e domain.Event) error
}

type PaymentAccount interface {
	UserUUID() uuid.UUID
	AccountUUID() uuid.UUID
	Status() domain.AccountStatus
	Balance() domain.Money

	Deposit(m domain.Money)
	Withdraw(m domain.Money) bool
	CanWithdraw(m domain.Money) bool
	SetStatus(s domain.AccountStatus)
}
