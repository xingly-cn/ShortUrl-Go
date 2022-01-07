package utils

import "regexp"

func UrlCheck(url string) bool {
	pattern := "^(((ht|f)tps?):\\/\\/)?[\\w-]+(\\.[\\w-]+)+([\\w.,@?^=%&:/~+#-]*[\\w@?^=%&/~+#-])?$"
	flag, _ := regexp.Match(pattern, []byte(url))
	return flag
}
