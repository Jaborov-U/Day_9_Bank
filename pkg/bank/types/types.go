package types

// Money = денежная сумма в минимальных единицах
type Money int64

// Currency = представляет код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN = номер карты
type PAN string

// Card = информация о платежной карте
type Card struct{
	ID int
	PAN PAN
	Balance Money
	Currency Currency
	Color string
	Name string
	Active bool
	MinBalance Money
}

// Payment пресдавляет информацию о платеже
type Payment struct {
	ID int
	Amount Money
}

// Слайс из карт для ДЗ 9
type PaymentSourse struct{
	Type string // card
	Number string // номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}