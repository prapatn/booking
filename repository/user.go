package repository

import (
	"booking/database"
	"booking/entities/users"
)

func CreateUser(user *users.Register) error {
	return database.DB.Table(user.TableName()).Create(user).Error
}

func UpdateUser(user interface{}, id string) int64 {
	return database.DB.Table("users").Where("id = ?", id).Updates(user).RowsAffected
}
