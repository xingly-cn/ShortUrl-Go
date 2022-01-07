package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// 获取数据库对象
func GetSQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:XNXxnx520@(8.142.199.134:3306)/shorturl?charset=utf8mb4")
	if err != nil {
		log.Println("连接数据库异常：", err)
	}
	return db
}
