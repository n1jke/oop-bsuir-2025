package domain

import "fmt"

// Item - товар в заказе
type Item struct {
	ID    string
	Name  string
	Price float64
}

// Address - адрес доставки
type Address struct {
	City   string
	Street string
	Zip    string
}

// todo : implement discount logic
// Order - заказ
type Order struct {
	ID          string
	Items       []Item
	Type        string // "Standard", "Premium", "Budget", "International"
	ClientEmail string
	Destination Address
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
