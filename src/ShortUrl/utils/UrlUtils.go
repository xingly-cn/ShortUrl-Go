package utils

import "regexp"

// Url 格式检查
func UrlCheck(url string) (bool, error) {
	pattern := "^(((ht|f)tps?):\\/\\/)?[\\w-]+(\\.[\\w-]+)+([\\w.,@?^=%&:/~+#-]*[\\w@?^=%&/~+#-])?$"
	flag, err := regexp.Match(pattern, []byte(url))
	return flag, err
}
