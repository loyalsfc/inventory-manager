package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/utils"
	"gorm.io/gorm"
)

type CategoryModel struct {
	DB *gorm.DB
}

type Category struct {
	ID   uuid.UUID `json:"id" gorm:"column:id;unique;not null"`
	Name string    `json:"name" gorm:"column:name;unique;not null"`
	Slug string    `json:"slug" gorm:"column:slug;unique;not null"`
	gorm.Model
}

func (c CategoryModel) CreateCategory(name string) (*Category, error) {
	if len(name) < 3 {
		return nil, errors.New("category name cannot be less than 3")
	}

	if categoryExist := c.IsCategoryExist(name); !categoryExist {
		return nil, errors.New("category already exist")
	}

	category := &Category{
		Name: name,
		Slug: utils.GenerateSlugs(name),
		ID:   uuid.New(),
	}

	if result := c.DB.Create(&category); result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (c CategoryModel) IsCategoryExist(name string) bool {
	result := c.DB.Where("name = ?", name).First(&Category{})

	return result.Error != nil
}

func (c CategoryModel) FindCategoryById(id uuid.UUID) (*Category, error) {
	var category Category
	result := c.DB.Where("id = ?", id).First(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (c CategoryModel) DeleteCategory(id uuid.UUID) error {
	category, err := c.FindCategoryById(id)

	if err != nil {
		return err
	}

	result := c.DB.Delete(&category)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c CategoryModel) EditCategory(name string, catID uuid.UUID) error {
	if len(name) < 3 {
		return errors.New("category name cannot be less than 3")
	}

	if categoryExist := c.IsCategoryExist(name); !categoryExist {
		return errors.New("category does not exist")
	}

	result := c.DB.Save(Category{
		ID:   catID,
		Name: name,
		Slug: utils.GenerateSlugs(name),
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c CategoryModel) CategoryList() ([]Category, error) {
	var categories []Category

	result := c.DB.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}
