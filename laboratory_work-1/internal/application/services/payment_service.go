package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

var (
	ErrAccountNotActive = errors.New("account is not active")
	ErrWithdrawRejected = errors.New("withdraw rejected")
	ErrNotImplemented   = errors.New("not implemented")
)

type PaymentService struct {
	accounts application.AccountRepository
	events   application.EventService
}

func NewPaymentService(accounts application.AccountRepository, events application.EventService) *PaymentService {
	return &PaymentService{accounts: accounts, events: events}
}

func (s *PaymentService) Deposit(m domain.Money, aUUID uuid.UUID) error {
	account, err := s.accounts.ByUUID(aUUID)
	if err != nil {
		return err
	}

	if account.Status() != domain.Active {
		return ErrAccountNotActive
	}

	account.Deposit(m)
	return nil
}

func (s *PaymentService) Withdraw(m domain.Money, aUUID uuid.UUID) error {
	account, err := s.accounts.ByUUID(aUUID)
	if err != nil {
		return err
	}

	if account.Status() != domain.Active {
		return ErrAccountNotActive
	}

	if !account.Withdraw(m) {
		return ErrWithdrawRejected
	}
	return nil
}

func (s *PaymentService) ProcessTransaction(t domain.Transaction) error {
	if t.Type() != domain.Transfer {
		return ErrNotImplemented
	}

	if err := s.Withdraw(t.Value(), t.SourceAccountUUID()); err != nil {
		return err
	}
	if err := s.Deposit(t.Value(), t.DestinationAccountUUID()); err != nil {
		_ = s.Deposit(t.Value(), t.SourceAccountUUID())
		return err
	}

	return nil
}
