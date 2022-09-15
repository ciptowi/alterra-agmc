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
func GetAll() (User, error) {
	var user User
	//result := DB.Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)
	if err := DB.Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetOneById(id int) (User, error) {
	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUserById(id int) error {
	var user User
	if err := DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int) error {
	var user User
	if err := DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
