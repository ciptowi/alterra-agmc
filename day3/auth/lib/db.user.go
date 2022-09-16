package lib

import (
	"auth/config"
	"auth/models"
)

var db, _ = config.Setup()

func CreateUser(user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func FindUsers(users *models.Users) error {
	if err := db.Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func FindUserById(id int, user *models.User) error {
	if err := db.First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByEmail(email string, user *models.User) error {
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByName(name string, user *models.User) error {
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserById(id int, user *models.User) error {
	if err := db.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int, user *models.User) error {
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
