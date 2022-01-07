package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 获取数据库对象
func GetSQL() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:XNXxnx520@(8.142.199.134:3306)/teapic?charset=utf8mb4")
	return db, err
}
