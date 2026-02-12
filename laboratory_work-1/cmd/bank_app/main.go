package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/application/services"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/domain"
	"github.com/n1jke/oop-bsuir-2025/lr-1/internal/infrastructure"
)

func main() {
	accountStore := infrastructure.NewMemoryAccountStorage()
	eventStore := infrastructure.NewMemoryEventStorage()

	accountRepo := services.NewAccountRepository(accountStore)
	eventService := services.NewEventService(eventStore)
	paymentService := services.NewPaymentService(accountRepo, eventService)
	useCase := application.NewTransferUseCase(accountRepo, paymentService, eventService)

	clientID := uuid.New()
	srcAccount := domain.NewCheckingAccount(clientID, uuid.New(), "BYN")
	destAccount := domain.NewSavingsAccount(clientID, uuid.New(), "BYN", domain.Silver)

	_ = accountStore.Save(srcAccount)
	_ = accountStore.Save(destAccount)

	_ = paymentService.Deposit(domain.NewMoney(2000, "BYN"), srcAccount.AccountUUID())

	tx := domain.NewTransaction(
		1,
		srcAccount.AccountUUID(),
		destAccount.AccountUUID(),
		domain.NewMoney(1000, "BYN"),
		domain.Transfer,
	)

	if err := useCase.Execute(&tx); err != nil {
		log.Printf("transfer failed: %v", err)
		return
	}

	events := eventService.QueryAll()
	log.Printf("transfer completed, events: %d", len(events))
	for _, e := range events {
		log.Printf("event: %s, account: %s, occurred at: %s", e.Name, e.AccountUUID, e.OccurredAt)
	}
}
