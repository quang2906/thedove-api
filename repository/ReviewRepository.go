package repository

import (
	"errors"
	"time"

	"ocg.com/database"
	"ocg.com/model"
)

var reviews []*model.Review

func CreateNewReview(review *model.Review) int64 {
	review.CreatedAt = time.Now().Unix()
	review.ModifiedAt = time.Now().Unix()
	database.DB.Create(&review)
	return review.Id
}

func GetAllReviews() []*model.Review {
	database.DB.Find(reviews)
	return reviews
}

func FindReviewById(Id int64) (*model.Review, error) {
	review := new(model.Review)
	database.DB.Where("id = ?", Id).Find(&review)
	if review != nil {
		return review, nil
	} else {
		return nil, errors.New("review not found")
	}
}

func FindReviewById2(Id int64) *model.Review {
	review := new(model.Review)
	database.DB.Where("id = ?", Id).Find(&review)
	if review != nil {
		return review
	} else {
		return nil
	}
}

func DeleteReviewById(Id int64) error {
	review := new(model.Review)
	database.DB.Where("id = ?", Id).Delete(&review)
	if review != nil {
		return nil
	} else {
		return errors.New("review not found")
	}
}

func UpdateReviewById(Id int64, reviewUpdate *model.Review) error {
	review := new(model.Review)
	database.DB.Where("id = ?", Id).Find(&review)
	if review != nil {
		review = reviewUpdate
		database.DB.Save(&review)
		return nil
	} else {
		return errors.New("review not found")
	}
}

func UpsertReviewById(Id int64, reviewUpsert *model.Review) {
	review := new(model.Review)
	database.DB.Where("id = ?", Id).Find(&review)
	if review != nil {
		review = reviewUpsert
		database.DB.Save(&review)
	} else {
		CreateNewReview(review)
	}
}
