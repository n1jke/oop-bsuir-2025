package main

import "fmt"

// =========================================================
// Файл: staff.go
// Описание: Система управления персоналом склада.
// =========================================================

type WarehouseWorker interface {
	ProcessOrder()
	AttendMeeting()
	GetRest()
	SwingingTheLead()
}

// HumanManager - Человек
type HumanManager struct{}

func (h HumanManager) ProcessOrder() {
	fmt.Println("Manager is processing logic...")
}

func (h HumanManager) AttendMeeting() {
	fmt.Println("Manager is boring at the meeting...")
}

func (h HumanManager) GetRest() {
	fmt.Println("Manager is taking a break...")
}

func (h HumanManager) SwingingTheLead() {
	fmt.Println("Manager is watching reels...")
}

// RobotPacker - Робот
type RobotPacker struct {
	Model string
}

func (r RobotPacker) ProcessOrder() {
	fmt.Println("Robot " + r.Model + " is packing boxes...")
}

func (r RobotPacker) GetRest() {
	fmt.Println("Robot was taken for maintenance")
}

func (r RobotPacker) AttendMeeting() {
	fmt.Println("ERROR: Robot cannot attend meetings")
}

func (r RobotPacker) SwingingTheLead() {
	panic("CRITICAL ERROR: Robot cannot waste our money (we hope so)")
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
