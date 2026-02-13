package main

import (
	"fmt"
	"log"
)

// =========================================================
// Файл: main.go
// Описание: Точка входа в приложение.
// =========================================================

func main() {
	// 1. Создание заказа
	order := Order{
		ID:   "ORD-256-X",
		Type: "Premium",
		Items: []Item{
			{ID: "1", Name: "Thermal Clips", Price: 1500},
			{ID: "2", Name: "UNATCO Pass Card", Price: 50},
		},
		ClientEmail: "jeevacation@gmail.com",
		Destination: Address{City: "Agartha", Street: "33 Thomas Street", Zip: "[REDACTED]"},
	}

	// 2. Инициализация процессора
	processor := NewOrderProcessor()

	// 3. Обработка заказа
	if err := processor.Process(order); err != nil {
		log.Fatalf("Failed to process order: %v", err)
	}

	// 4. Работа с обслуживанием
	fmt.Println("\nTesting Warehouse Stuff:")
	workers := []WarehouseWorker{
		HumanManager{},
		RobotPacker{Model: "George Droid"},
	}

	ManageWarehouse(workers)
}
