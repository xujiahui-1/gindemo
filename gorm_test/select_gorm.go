package main

import "fmt"

type UserInfo struct {
	Name string
	Age  uint8
}

func SelectTset() {
	var u []UserInfo
	GLOBAL_DB.Model(&User{}).Find(&u)
	for _, v := range u {
		fmt.Println(v)
	}

}
