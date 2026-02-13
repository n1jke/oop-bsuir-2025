package main

// =========================================================
// Файл: models.go
// Описание: Модели данных системы.
// =========================================================

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

// Order - заказ
type Order struct {
	ID          string
	Items       []Item
	Type        string // "Standard", "Premium", "Budget", "International"
	ClientEmail string
	Destination Address
}
