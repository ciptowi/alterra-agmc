package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *gorm.DB

func CreateUser(usr interface{}) error {
	if err := DB.Create(&usr).Error; err != nil {
		return err
	}
	return nil
}

//func GetAll(keywords string) ([]User, error) {
func GetAll(usr interface{}) error {
	//result := DB.Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)
	if err := DB.Find(&usr).Error; err != nil {
		return err
	}
	return nil
}

func GetOneById(id int) (User, error) {
	var user User
	result := DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func UpdateUserById(id int) (User, error) {
	var user User
	result := DB.Where("id = ?", id).Updates(&user)
	return user, result.Error
}

func DeleteUserById(id int) (User, error) {
	var user User
	result := DB.Where("id = ?", id).Delete(&user)
	return user, result.Error
}
