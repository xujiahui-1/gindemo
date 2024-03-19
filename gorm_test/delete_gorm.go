package main

import "fmt"

func DeleteTest() {
	var myUser User
	//GLOBAL_DB.Where("name=?", "qm").Delete(&myUser)            //逻辑删除
	//GLOBAL_DB.Unscoped().Where("name=?", "qm").Delete(&myUser) //Unscoped()直接删除
	GLOBAL_DB.Raw("SELECT * FROM t_user WHERE name = ?", "王晓懿").Scan(&myUser) //原生sql
	fmt.Println(myUser)

}
