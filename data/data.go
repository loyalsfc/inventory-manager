package data

import (
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/utils"
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

type FormData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserParams struct {
	Email string         `json:"email"`
	Role  utils.UserRole `json:"role"`
}
