package repository

import (
	"errors"

	"ocg.com/database"
	"ocg.com/model"
)

var categories []*model.Category

func CreateNewCategory(category *model.Category) int64 {
	database.DB.Create(&category)
	return category.Id
}

func GetAllCategories() []*model.Category {
	database.DB.Find(&categories)
	return categories
}

func FindCategoryById(Id int) (*model.Category, error) {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		return category, nil
	} else {
		return nil, errors.New("category not found")
	}
}

func FindCategoryById2(Id int64) *model.Category {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		return category
	} else {
		return nil
	}
}

func DeleteCategoryById(Id int64) error {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Delete(&category)
	if category != nil {
		return nil
	} else {
		return errors.New("category not found")
	}
}

func UpdateCategoryById(Id int64, categoryUpdate *model.Category) error {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		category = categoryUpdate
		database.DB.Save(&category)
		return nil
	} else {
		return errors.New("category not found")
	}
}

func UpsertCategoryById(Id int64, categoryUpsert *model.Category) {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		category = categoryUpsert
		database.DB.Save(&category)
	} else {
		CreateNewCategory(category)
	}
}
