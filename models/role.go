package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleName string `json:"role_name"`
}

const (
	RoleUser  = "ROLE_USER"
	RoleAdmin = "ROLE_ADMIN"
)
