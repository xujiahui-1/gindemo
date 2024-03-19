package main

func UpDateTest() {
	//update 只更新你选择的字段
	//GLOBAL_DB.Model(&User{}).Where("name =?", "徐佳辉").Update("name", "徐大帅")
	//updates 更新所有字段 一种为map,一种为结构体,结构体0值不参与更新
	//结构体
	var myUser User
	//GLOBAL_DB.First(&myUser).Updates(User{Name: "嘘嘘嘘", Age: 36}) //结构体0值不参与更新
	GLOBAL_DB.First(&myUser).Updates(map[string]interface{}{"Name": "王八蛋", "Age": 0}) //map的0值参与更新

	//Map
	////save 无论如何都更新 包括0值
	//var myUser []User
	//ddd := GLOBAL_DB.Model(&User{}).Where("name =?", "徐大帅").Find(&myUser)
	//for K := range myUser {
	//	myUser[K].Age = 18
	//}
	//ddd.Save(&myUser)
}
