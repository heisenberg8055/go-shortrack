package handlers

import (
	"crypto/md5"
	"encoding/hex"
)

func convertToBase62(number int) string {
	const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62 := ""
	for number > 0 {
		remainder := number % 62
		base62 = string(base62Digits[remainder]) + base62
		number /= 62
	}
	return base62
}

func convertToMD5(longURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(longURL))
	return hex.EncodeToString(hasher.Sum(nil))[:8]
}
