package services

import (
	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

type AccountStorage interface {
	Save(account application.PaymentAccount) error
	ByUUID(accountUUID uuid.UUID) (application.PaymentAccount, error)
	UpdateStatus(accountUUID uuid.UUID, status domain.AccountStatus) error
}

type EventStorage interface {
	Save(event domain.Event) error
	QueryAll() []domain.Event
}
