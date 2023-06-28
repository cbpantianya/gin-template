package utils

import (
	"math/rand"
	"time"
)

// AddRandomPrefix 给传入的字符串添加前缀
func AddRandomPrefix(str string) string {
	var source = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// 从其中随机选取8位随机字符
	var random string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		random += string(source[rand.Intn(len(source))])
	}
	return string(random) + "_" + str
}
