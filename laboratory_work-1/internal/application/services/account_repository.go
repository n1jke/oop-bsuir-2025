package services

import (
	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

type AccountRepository struct {
	storage AccountStorage
}

func NewAccountRepository(storage AccountStorage) *AccountRepository {
	return &AccountRepository{storage: storage}
}

func (r *AccountRepository) Create(account application.PaymentAccount) error {
	return r.storage.Save(account)
}

func (r *AccountRepository) ChangeStatus(accountUUID uuid.UUID, status domain.AccountStatus) error {
	return r.storage.UpdateStatus(accountUUID, status)
}

func (r *AccountRepository) ByUUID(accountUUID uuid.UUID) (application.PaymentAccount, error) {
	return r.storage.ByUUID(accountUUID)
}
