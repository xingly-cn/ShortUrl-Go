package utils

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetSQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:XNXxnx520@(8.142.199.134:3306)/shorturl?charset=utf8mb4")
	if err != nil {
		log.Println("连接数据库异常：", err)
	}
	return db
}

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "175.27.243.243:6379",
		Password: "213879",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println("Redis异常：", err)
	}
	return rdb
}
