package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"shorturl/ShortUrl/entity"
	"shorturl/ShortUrl/utils"
	"time"
)

var (
	rdb *redis.Client
	db  *gorm.DB
)

// 初始化
func init() {
	// 日志配置
	log.SetPrefix("【日志】")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 数据库启动
	db = utils.GetSQL()

	// Redis启动
	rdb = utils.GetRedis()
}

func add(c *gin.Context) {
	target := c.PostForm("url")
	if utils.UrlCheck(target) == false {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "URL格式错误,正确格式：【http(s)://suo.asugar.cn】",
		})
		return
	}

	id := utils.UrlEncode(target)

	err := rdb.Set(id, target, time.Duration(259200)*time.Second).Err()
	if err != nil {
		log.Println("Redis保存异常：", err)
	}

	db.Create(&entity.Urlinfo{
		Surl:       target,
		Lurl:       id,
		Views:      0,
		Createtime: time.Now(),
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "生成成功",
		"data": id,
	})
}

func visit(c *gin.Context) {
	id := c.Param("hash")
	url, err := rdb.Get(id).Result()
	urlinfo := entity.Urlinfo{}

	// 如果 Redis 中找到
	if err == nil && len(url) > 1 {
		c.Redirect(http.StatusFound, url)
		// 链接被访问,更新过期时间
		rdb.Expire(id, time.Duration(259200)*time.Second)
		// 更新访问量
		db.Where("lurl = ?", id).First(&urlinfo)
		if urlinfo != (entity.Urlinfo{}) {
			urlinfo.Views++
			db.Model(&urlinfo).Update("views", urlinfo.Views)
		}
		return
	}

	// 如果 MySQL 中找到
	db.Where("lurl = ?", id).First(&urlinfo)
	if len(urlinfo.Surl) > 1 {
		c.Redirect(http.StatusFound, urlinfo.Surl)
		// 存入 Redis
		err := rdb.Set(id, urlinfo.Surl, time.Duration(259200)*time.Second).Err()
		if urlinfo != (entity.Urlinfo{}) {
			urlinfo.Views++
			db.Model(&urlinfo).Update("views", urlinfo.Views)
		}
		if err != nil {
			log.Println("Redis保存异常：", err)
		}
		return
	}

	// 不存在直接返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Not Found",
	})
}

func main() {

	r := gin.Default()


	// 用户功能
	r.POST("/make", add)
	r.GET("/go/:hash", visit)
	
	// 管理功能 - 未完待续

	// 404 处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "页面不存在"})
	})

	r.Run()
}
