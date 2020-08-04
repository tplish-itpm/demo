package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tplish-itpm/demo/pkg/e"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
	Role     Role   `json:"-"`
}

//func FindUserById(id int64) (*User, e.Error) {
//	var user User
//	err := DB.Where("id = ?", id).First(&user).Error
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return nil, nil
//		}
//		return nil, e.ErrNewErr(err)
//	}
//	return &user, nil
//}

func FindUserByUsername(username string) (*User, e.Error) {
	var user User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, e.ErrNewErr(err)
	}
	return &user, nil
}

func FindUserByUsernameWithRole(username string) (*User, e.Error) {
	var user User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, e.ErrNewErr(err)
	}
	DB.Where("id = ?", user.RoleId).Find(&user.Role)

	return &user, nil
}

func ListUsers() ([]User, e.Error) {
	var users []User
	err := DB.Find(&users).Error
	return users, e.ErrNewErr(err)
}

func AddUser(user *User) e.Error {
	var role Role
	if err := DB.Where("role_name = ?", RoleUser).First(&role).Error; err != nil {
		return e.ErrNewErr(err)
	}
	user.Role = role
	if err := DB.Create(user).Error; err != nil {
		return e.ErrNewErr(err)
	}
	return nil
}
