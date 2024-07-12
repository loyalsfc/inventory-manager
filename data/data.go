package data

import (
	"github.com/google/uuid"
)

type AddProductParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	CategoryId  uuid.UUID `json:"category_id"`
}

type PaymentMethod struct {
	Cash     int `json:"cash"`
	Transfer int `json:"transfer"`
	Pos      int `json:"pos"`
}

type OrderProducts struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type OrderParams struct {
	Products      []OrderProducts `json:"products"`
	PaymentMethod PaymentMethod   `json:"payment_method"`
}
