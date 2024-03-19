package main

import "gorm.io/gorm"

/*
https://gorm.io/zh_CN/docs/models.html

	这样一个空结构体已经是一个gorm的模型了
*/
type User struct {
	gorm.Model
	Name string `gorm:"default:qm"`
	Age  uint8
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&User{})
}
