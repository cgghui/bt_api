package cgghui

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString 随机字符串 数字+大小字母
func RandomString(lenX int, strL ...string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if len(strL) == 1 {
		str = strL[0]
	}
	size := len(str)
	byteArr := []byte(str)
	var result []byte
	for i := 0; i < lenX; i++ {
		result = append(result, byteArr[rand.Intn(size)])
	}
	return string(result)
}

// RandomSliceString 从切片中随机返回
func RandomSliceString(arr *[]string) string {
	return (*arr)[rand.Intn(len(*arr))]
}

// MD5Byte MD5值
func MD5Byte(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 MD5值
func MD5(str string) string {
	return MD5Byte([]byte(str))
}

// Str2Int string转int
func Str2Int(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = def
	}
	return i
}
