package repository

import (
	"errors"

	"ocg.com/database"
	"ocg.com/model"
)

var users []*model.User

func CreateNewUser(user *model.User) int64 {
	database.DB.Create(&users)
	return user.Id
}

func GetAllUsers() []*model.User {
	database.DB.Find(&reviews)
	return users
}

func FindUserById(Id int64) (*model.User, error) {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Find(&user)
	if user != nil {
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func FindUserById2(Id int64) *model.User {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Find(&user)
	if user != nil {
		return user
	} else {
		return nil
	}
}

func DeleteUserById(Id int64) error {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Delete(&user)
	if user != nil {
		return nil
	} else {
		return errors.New("user not found")
	}
}

func UpdateUserById(Id int64, userUpdate *model.User) error {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Find(&user)
	if user != nil {
		user = userUpdate
		database.DB.Save(&user)
		return nil
	} else {
		return errors.New("user not found")
	}
}

func UpsertUserById(Id int64, userUpsert *model.User) {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Find(&user)
	if user != nil {
		user = userUpsert
		database.DB.Save(&user)
	} else {
		CreateNewUser(user)
	}
}
