package main

type JiaZi struct {
	ID          uint
	Name        string
	XiaoFengChe []XiaoFengChe `gorm:"foreignKey:JiaZiName;references:Name"` //配置了外键，会自动根据这列匹配 ,使用references重写引用
}

type XiaoFengChe struct {
	ID        uint
	Name      string
	JiaZiName string
}

func pro() {
	GLOBAL_DB.AutoMigrate(&JiaZi{}, &XiaoFengChe{})
	//结果发现JiaZiName字段还是存了小风车的id
	GLOBAL_DB.Create(&JiaZi{Name: "小架子", XiaoFengChe: []XiaoFengChe{
		{Name: "小风车1"},
		{Name: "小风车2"},
	}})

}
