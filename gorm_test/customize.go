package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 自定义数据类型
func (c CInfo) Value() (driver.Value, error) {
	//driver.Value为要存入数据库的字段
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (c *CInfo) Scan(value interface{}) error {
	//把 myInfo CInfo当作json存入数据库,取出来的时候在渲染回CInfo
	str, ok := value.([]byte)
	if !ok {
		return errors.New("不配")
	}
	json.Unmarshal(str, c)
	return nil
}

type CInfo struct {
	Name string
	Age  int
}
type CUser struct {
	MyInfo CInfo
	ID     uint
}

func Commm() {
	//GLOBAL_DB.AutoMigrate(&CUser{})
	GLOBAL_DB.Create(&CUser{ID: 1, MyInfo: CInfo{Name: "徐佳辉", Age: 18}})

	var sss CUser
	GLOBAL_DB.Find(&sss)
	fmt.Println(sss)
}
