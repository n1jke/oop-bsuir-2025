package application

import (
	"errors"
	"time"

	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
)

var ErrUnsupportedTransactionType = errors.New("unsupported transaction type")

type TransferUseCase struct {
	accounts AccountRepository
	payments PaymentService
	events   EventService
}

func NewTransferUseCase(
	accounts AccountRepository,
	payments PaymentService,
	events EventService,
) *TransferUseCase {
	return &TransferUseCase{
		accounts: accounts,
		payments: payments,
		events:   events,
	}
}

func (uc *TransferUseCase) Execute(t *domain.Transaction) error {
	if t.Type() != domain.Transfer {
		return ErrUnsupportedTransactionType
	}

	t.SetStatus(domain.Pending)

	total := t.Value()

	if err := uc.payments.Withdraw(total, t.SourceAccountUUID()); err != nil {
		t.SetStatus(domain.Declined)
		return err
	}

	if err := uc.payments.Deposit(t.Value(), t.DestinationAccountUUID()); err != nil {
		_ = uc.payments.Deposit(total, t.SourceAccountUUID())
		t.SetStatus(domain.Failed)
		return err
	}

	t.SetStatus(domain.Completed)

	if uc.events != nil {
		now := time.Now()
		_ = uc.events.Publish(domain.Event{
			Name:        "transfer_completed_src",
			AccountUUID: t.SourceAccountUUID(),
			OccurredAt:  now,
		})
		_ = uc.events.Publish(domain.Event{
			Name:        "transfer_completed_dest",
			AccountUUID: t.DestinationAccountUUID(),
			OccurredAt:  now,
		})
	}

	return nil
}
