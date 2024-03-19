package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money int
	DogID uint
}
type Dog struct {
	gorm.Model
	Name     string
	Infoo    Info
	GirlGods []GirlGod `gorm:"many2many:dog_girl_god"`
}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god"`
}

func One2One() {
	myDog := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	//Preload自动从中间表中找到关联数据
	GLOBAL_DB.Preload("GirlGods").Find(&myDog)
	//当我们只要关联表数据而不要原表数据的时候
	//声明一个承载结果的结构体
	var girls []GirlGod
	//Association关联模式关联Dog里的女神数组
	//GLOBAL_DB.Model(&myDog).Association("GirlGods").Find(&girls)
	//查询出来女神的数据还想把狗也带出来
	GLOBAL_DB.Model(&myDog).Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Infoo").Where("money>?", "1000")
	}).Association("GirlGods").Find(&girls)
	fmt.Println(girls)
}
