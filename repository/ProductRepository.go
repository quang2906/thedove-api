package repository

import (
	"errors"
	"time"

	"ocg.com/database"
	"ocg.com/model"
)

var products []*model.Product

func CreateNewProduct(product *model.Product) int64 {
	product.CreatedAt = time.Now().Unix()
	product.ModifiedAt = time.Now().Unix()
	database.DB.Create(&product)
	return product.Id
}

func GetAllProducts() []*model.Product {
	database.DB.Preload("Image").Find(&products)
	return products
}

func FindProductById(Id int64) (*model.Product, error) {
	product := new(model.Product)
	database.DB.Preload("Image").Where("id = ?", Id).Find(&product)
	if product != nil {
		return product, nil
	}
	return nil, errors.New("product not found")
}

func FindProductById2(Id int64) *model.Product {
	product := new(model.Product)
	database.DB.Where("id = ?", Id).Find(&product)
	return product
}

func DeleteProductById(Id int64) error {
	product := new(model.Product)
	database.DB.Where("id = ?", Id).Delete(&product)
	if product != nil {
		return nil
	}
	return errors.New("product not found")
}

func UpdateProductById(Id int64, productUpdate *model.Product) error {
	product := new(model.Product)
	database.DB.Where("id = ?", Id).Find(&product)
	if product != nil {
		product = productUpdate
		database.DB.Save(&product)
		return nil
	}
	return errors.New("product not found")
}

func UpsertProductById(Id int64, productUpsert *model.Product) {
	product := new(model.Product)
	database.DB.Where("id = ?", Id).Find(&product)
	if product != nil {
		product = productUpsert
		database.DB.Save(&product)
	} else {
		CreateNewProduct(product)
	}
}

// func UpdateProductRating(product *model.Product) *model.Product {
// 	var sumRating float64 = 0
// 	countRating := 0
// 	reviews := Reviews.GetAllReviews()
// 	for _, review := range reviews {
// 		if review.ProductId == product.Id {
// 			countRating++
// 			sumRating += review.Rating
// 		}
// 	}
// 	product.Rating = float64(sumRating) / float64(countRating)
// 	return product
// }
