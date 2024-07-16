package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/data"
	"gorm.io/gorm"
)

type Status string

const (
	pending   Status = "pending"
	completed Status = "completed"
	failed    Status = "failed"
)

type Order struct {
	OrderID       uuid.UUID          `json:"order_id" gorm:"column:order_id;unique;primary;not null"`
	Status        Status             `json:"status" gorm:"column:status;not null"`
	Product       json.RawMessage    `json:"products" gorm:"foreignKey:product_id;column:products;type:jsonb;not null"`
	PaymentMethod data.PaymentMethod `json:"payment_method" gorm:"column:payment_method;embedded;not null"`
	TotalPrice    int                `json:"total_price" gorm:"column:total_price;not null"`
	gorm.Model
}

type OrderService struct {
	DB *gorm.DB
}

func (o OrderService) CreateOrder(param data.OrderParams) (*Order, error) {
	//calculate the amount paid
	amountPaid := param.PaymentMethod.Cash + param.PaymentMethod.Pos + param.PaymentMethod.Transfer

	if amountPaid == 0 {
		return nil, errors.New("please include a valid payment method")
	}

	// Calculate total order price
	totalPrice := 0

	//Check if all products are available and the quantity required
	for _, product := range param.Products {
		if product.Quantity < 1 {
			return nil, fmt.Errorf("invalid quantity for product with id %v", product.ProductID)
		}

		var item Product
		result := o.DB.Where("id = ?", product.ProductID).First(&item)
		if result.Error != nil {
			return nil, fmt.Errorf("no product found for id %v", product.ProductID)
		}

		totalPrice = totalPrice + (item.Price * product.Quantity)
	}

	if totalPrice != amountPaid {
		return nil, errors.New("total amount does not match")
	}

	marshal, err := json.Marshal(param.Products)

	if err != nil {
		return nil, err
	}

	order := Order{
		OrderID:       uuid.New(),
		Status:        completed,
		Product:       marshal,
		PaymentMethod: param.PaymentMethod,
		TotalPrice:    totalPrice,
	}

	if result := o.DB.Create(&order); result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o OrderService) GetAllOrders() ([]Order, error) {
	var orders []Order

	if result := o.DB.Find(&orders); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o OrderService) FindOrder(id uuid.UUID) (*Order, error) {
	var order Order
	if result := o.DB.Where("order_id = ?", id).First(&order); result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o OrderService) DeleteOrder(id uuid.UUID) error {
	if result := o.DB.Where("order_id = ?", id).Delete(&Order{}); result.Error != nil {
		return result.Error
	}

	return nil
}
