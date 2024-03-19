package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 声明一个全局变量来拿到gorm.DB的指针
var GLOBAL_DB *gorm.DB

func main() {

	db, _ := gorm.Open(mysql.New(mysql.Config{
		//dsn建立连接
		DSN:                       "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         171,   //否则以字符串为主键时会造成索引超长
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 功能上的详细配置
		SkipDefaultTransaction: false, //是否跳过默认事务
		//是否复数表名，就是以结构体等创建表的时候增加前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", //前缀
			SingularTable: true, //true为单数表名，加上后表后就加s了

		},
		//ture后建表的时候就不会设置物理外键，现在主张逻辑外键
		DisableForeignKeyConstraintWhenMigrating: true, //创建表的时候是否跳过外键约束的建立
	})

	/**
	数据库连接池配置
	*/
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           //连接池最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          //连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //连接池中连接的最大可复用时间

	//将数据库指针赋给全局变量GLOBAL_DB，方便别的包用
	GLOBAL_DB = db

	//调用建表方法
	//TestUserCreate()
	//调用插入方法
	//CreatedTest()
	//调用查询方法
	//SelectTset()
	//调用修改方法
	//UpDateTest()
	//删除方法
	//DeleteTest()
	//One2One()
	//多态
	//pro()
	//事务
	//TransactionTest()
	//自定义数据类型
	Commm()
}
