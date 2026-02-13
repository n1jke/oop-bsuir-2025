package application

import "fmt"

// =========================================================
// Файл: staff.go
// Описание: Система управления персоналом склада.
// =========================================================

// todo define to some interfaces
type WarehouseWorker interface {
	ProcessOrder()
	AttendMeeting()
	GetRest()
	SwingingTheLead()
}

// ManageWarehouse - функция, которая работает со списком работников
func ManageWarehouse(workers []WarehouseWorker) {
	fmt.Println("\n--- Warehouse Shift Started ---")
	for _, worker := range workers {
		worker.ProcessOrder()
		worker.AttendMeeting()
		worker.GetRest()
		worker.SwingingTheLead()
	}
}
