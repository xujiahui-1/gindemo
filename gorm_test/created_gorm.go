package main

import "fmt"

func CreatedTest() {
	dbres := GLOBAL_DB.Create(&User{Name: "王晓懿", Age: 22})
	GLOBAL_DB.Select("Name").Create(&User{Name: "王晓懿"}) //选择创建,只想插入name
	GLOBAL_DB.Omit("Name").Create(&User{Age: 22})       //除了Name以外的字段都创建
	/**
	批量插入
	*/
	GLOBAL_DB.Create(&[]User{
		{Name: "王晓懿", Age: 22},
		{Name: "王晓懿", Age: 23},
		{Name: "王晓懿", Age: 2},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected) //是否报错和影响的条数
}
