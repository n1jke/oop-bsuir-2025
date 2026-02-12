package infrastructure

import (
	"errors"

	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

var ErrNotFound = errors.New("not found")

type MemoryAccountStorage struct {
	byID map[uuid.UUID]application.PaymentAccount
}

func NewMemoryAccountStorage() *MemoryAccountStorage {
	return &MemoryAccountStorage{byID: make(map[uuid.UUID]application.PaymentAccount)}
}

func (s *MemoryAccountStorage) Save(account application.PaymentAccount) error {
	s.byID[account.AccountUUID()] = account
	return nil
}

func (s *MemoryAccountStorage) ByUUID(accountUUID uuid.UUID) (application.PaymentAccount, error) {
	account, exist := s.byID[accountUUID]
	if !exist {
		return nil, ErrNotFound
	}

	return account, nil
}

func (s *MemoryAccountStorage) UpdateStatus(accountUUID uuid.UUID, status domain.AccountStatus) error {
	account, exist := s.byID[accountUUID]
	if !exist {
		return ErrNotFound
	}
	account.SetStatus(status)
	s.byID[accountUUID] = account
	return nil
}

type MemoryEventStorage struct {
	events []domain.Event
}

func NewMemoryEventStorage() *MemoryEventStorage {
	return &MemoryEventStorage{
		events: make([]domain.Event, 0, 16),
	}
}

func (s *MemoryEventStorage) Save(event domain.Event) error {
	s.events = append(s.events, event)
	return nil
}

func (s *MemoryEventStorage) QueryAll() []domain.Event {
	return s.events
}
