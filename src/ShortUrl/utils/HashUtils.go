package utils

import (
	"log"
	"math"
)

const (
	BIG_M = 0xc6a4a7935bd1e995
	BIG_R = 47
	SEED  = 0x1234ABCD
)

func MurmurHash64A(data []byte) (h int64) {
	var k int64
	h = SEED ^ int64(uint64(len(data))*BIG_M)

	var ubigm uint64 = BIG_M
	var ibigm = int64(ubigm)
	for l := len(data); l >= 8; l -= 8 {
		k = int64(int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 |
			int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56)

		k = k * ibigm
		k ^= int64(uint64(k) >> BIG_R)
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch len(data) {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= ibigm
	}

	h ^= int64(uint64(h) >> BIG_R)
	h *= ibigm
	h ^= int64(uint64(h) >> BIG_R)
	if h < 0 {
		h = math.MaxInt64 + h
	}
	return
}

const CODE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CodeLength = 62

func Base62Encode(number int64) string {
	if number == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for number > 0 {
		round := number / CodeLength
		remain := number % CodeLength
		result = append(result, CODE62[remain])
		number = round
	}
	return string(result)
}

func UrlEncode(url string) (res string) {
	t := MurmurHash64A([]byte(url))
	res = Base62Encode(t)
	if res == "" {
		log.Println("[" + url + "]" + "生成失败,请更换哈希函数！")
	}
	return
}
