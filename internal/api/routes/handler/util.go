package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func convertToBase62(number int64) string {
	const base62Digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base62 := ""
	for number > 0 {
		remainder := number % 62
		base62 = string(base62Digits[remainder]) + base62
		number /= 62
	}
	return base62
}

func convertToHash(longURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(longURL))
	md5Hash := hex.EncodeToString(hasher.Sum(nil))[:12]
	md5Int, _ := strconv.ParseInt(md5Hash, 16, 64)
	return convertToBase62(md5Int)
}
