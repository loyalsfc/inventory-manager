package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/data"
	"github.com/loyalsfc/investrite/utils"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"column:id;primarykey;not null;unique"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Description string    `json:"description" gorm:"column:description"`
	Quantity    int       `json:"quantity" gorm:"column:quantity;default:0;check=>0;not null"`
	Price       int       `json:"price" gorm:"column:price;not null"`
	Image       string    `json:"image" gorm:"column:image;"`
	CategoryId  uuid.UUID `json:"category_id" gorm:"column:category_id;not null"`
	Slug        string    `json:"slug" gorm:"column:slug;not null;unique"`
	gorm.Model
}

type ProductService struct {
	DB *gorm.DB
}

func (p ProductService) IsProductExist(id uuid.UUID) bool {
	var product = &Product{ID: id}
	if result := p.DB.First(&product); result.Error == nil {
		return true
	}
	return false
}

func (p ProductService) GetProductById(id uuid.UUID) (product *Product, err error) {
	var item Product

	result := p.DB.Where("id = ?", id).First(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (p ProductService) CreateProduct(data *data.AddProductParams) (*Product, error) {
	if len(data.Name) < 3 {
		return nil, errors.New("invalid product name")
	}

	var category = Category{}
	if result := p.DB.Where("id = ?", data.CategoryId).First(&category); result.Error != nil {
		return nil, errors.New("category id does not exist")
	}

	product := Product{
		ID:          uuid.New(),
		Name:        data.Name,
		Description: data.Description,
		Quantity:    data.Quantity,
		Price:       data.Price,
		Image:       data.Image,
		CategoryId:  data.CategoryId,
		Slug:        utils.GenerateSlugs(data.Name),
	}

	if result := p.DB.Create(&product); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (p ProductService) GetAllProducts() (*[]Product, error) {
	var products []Product

	if result := p.DB.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (p ProductService) UpdateProduct(id uuid.UUID, data *data.AddProductParams) error {
	product, err := p.GetProductById(id)
	if err != nil {
		return err
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Quantity = data.Quantity
	product.Price = data.Price
	product.Image = data.Image
	product.Slug = utils.GenerateSlugs(data.Name)

	if result := p.DB.Save(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p ProductService) DeleteProduct(id uuid.UUID) error {
	product, err := p.GetProductById(id)

	if err != nil {
		return err
	}

	if result := p.DB.Where("id = ?", id).Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p ProductService) IncrementQuantity(id uuid.UUID) (int, error) {
	product, err := p.GetProductById(id)

	if err != nil {
		return 0, err
	}

	product.Quantity = product.Quantity + 1

	if result := p.DB.Save(&product); result.Error != nil {
		return 0, err
	}

	return product.Quantity, nil
}

func (p ProductService) DecreaseQuantity(id uuid.UUID) (int, error) {
	product, err := p.GetProductById(id)

	if err != nil {
		return 0, err
	}

	if product.Quantity == 0 {
		return product.Quantity, nil
	}

	product.Quantity = product.Quantity - 1

	if result := p.DB.Save(&product); result.Error != nil {
		return 0, err
	}

	return product.Quantity, nil
}
