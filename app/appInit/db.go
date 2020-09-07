package appInit

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	mysqlArgs := Config.DB.User + ":" + Config.DB.Password + "@tcp(" + Config.DB.Host + ":" + Config.DB.Port + ")/" + Config.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	//Db, err = gorm.Open("mysql", "root:123456789@/chat?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("mysql", mysqlArgs)

	//mysql连接池设置
	DB.SingularTable(true)                       //全局禁用表名复数
	DB.DB().SetMaxOpenConns(300)                 //最大连接数
	DB.DB().SetMaxIdleConns(100)                 //最大空闲连接数
	DB.DB().SetConnMaxLifetime(30 * time.Second) //每个连接的过期时间

	if err != nil {
		panic(err)
	}

	//defer DB.Close()
	fmt.Println("DB connect success!!!", mysqlArgs)
	DB.LogMode(true)
}
