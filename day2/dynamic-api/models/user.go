package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *gorm.DB

func CreateUser(usr *User) error {
	if e := DB.Create(usr).Error; e != nil {
		return e
	}
	return nil
}

//func GetAll(keywords string) ([]User, error) {
func GetAll() ([]User, error) {
	var users []User
	//result := DB.Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)
	result := DB.Find(&users)

	return users, result.Error
}

func GetOneById(id int) (User, error) {
	var user User
	result := DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (user *User) UpdateUser(email string) error {
	if err := DB.Model(&User{}).Where("email = ?", email).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) DeleteUser() error {
	if err := DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
