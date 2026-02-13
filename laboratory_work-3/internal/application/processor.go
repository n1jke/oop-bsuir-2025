package application

import (
	"errors"
	"fmt"

	"github.com/n1jke/oop-bsuir-2025/lr-3/internal/domain"
)

// =========================================================
// Файл: processor.go
// Описание: Основная бизнес-логика.
// =========================================================

type Repository interface {
	SaveOrder(order domain.Order, total float64) error
}

type ClientNotifier interface {
	Notify(to string, subject string, body string)
}

type ManagerNotifier interface {
	Notify(to string, subject string, body string)
}

type OrderProcessor struct {
	database   Repository
	clientMsg  ClientNotifier
	managerMsg ManagerNotifier
}

func NewOrderProcessor(db Repository, cMsg ClientNotifier, mMsg ManagerNotifier) *OrderProcessor {
	return &OrderProcessor{
		database:   db,
		clientMsg:  cMsg,
		managerMsg: mMsg,
	}
}

func (op *OrderProcessor) Process(order domain.Order) error {
	fmt.Printf("--- Processing Order %s ---\n", order.ID)

	// 1. Логика валидации
	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}
	if order.Destination.City == "" {
		return errors.New("destination city is required")
	}

	// 2. Логика расчета суммы
	var total float64
	for _, item := range order.Items {
		total += item.Price
	}

	// 3. Логика скидок и налогов
	switch order.Type {
	case "Standard":
		// Стандартный налог
		total = total * 1.2
	case "Premium":
		// Скидка 10% + налог
		total = (total * 0.9) * 1.2
	case "Budget":
		if len(order.Items) > 3 {
			fmt.Println("Budget orders cannot have more than 3 items. Skipping.")
			return nil
		}
	case "International":
		total = total * 1.5 // Таможенный сбор
		if order.Destination.City == "Nowhere" {
			return errors.New("cannot ship to Nowhere")
		}
	default:
		return errors.New("unknown order type")
	}

	// 4. Логика сохранения
	if err := op.database.SaveOrder(order, total); err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	// 5. Логика уведомлений
	// msgs
	emailBody := fmt.Sprintf("<h1>Your order %s is confirmed!</h1><p>Total: %.2f</p>", order.ID, total)
	telegramBody := fmt.Sprintf("<h1>Order for client %s with orderId %s is confirmed!</h1><p>Total: %.2f</p>", order.ClientEmail, order.ID, total)
	//notify
	op.clientMsg.Notify(order.ClientEmail, "Order Confirmation", emailBody)

	// todo
	op.managerMsg.Notify("manager", "Order Notification", telegramBody)

	return nil
}
