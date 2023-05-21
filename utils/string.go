package utils

import (
	"math/rand"
	"time"
)

// 给传入的字符串添加前缀
func AddRandomPrefix(str string) string {
	var souce = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// 从其中随机选取8位随机字符
	var random string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8 ; i++ {
		random += string(souce[rand.Intn(len(souce))])
	}
	return string(random) + "_" +str
}
