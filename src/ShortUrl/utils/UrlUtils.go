package utils

import (
	"log"
	"regexp"
)

func UrlCheck(url string) bool {
	pattern := "^(((ht|f)tps?):\\/\\/)?[\\w-]+(\\.[\\w-]+)+([\\w.,@?^=%&:/~+#-]*[\\w@?^=%&/~+#-])?$"
	flag, err := regexp.Match(pattern, []byte(url))
	if err != nil {
		log.Println("URL格式检查异常：", err)
	}
	return flag
}
