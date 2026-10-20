package goutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// Implode 将任意类型数组，按照分割符，转成字符串
func Implode[T any](arr []T, delimiter string) string {
	var result string
	for i, v := range arr {
		if i > 0 {
			result += delimiter
		}
		result += fmt.Sprintf("%v", v)
	}
	return result
}

// InSlice 判断一个元素，是否包含在一个切片内

func InSlice[T comparable](need T, arr []T) bool {
	for _, v := range arr {
		if v == need {
			return true
		}
	}
	return false
}

//Strpos 判断一个字符串在另一个字符串的首次出现的位置

func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}
func Ucfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	firstRune := []rune(s)[0]
	if unicode.IsLower(firstRune) {
		firstRune = unicode.ToUpper(firstRune)
	}
	return string(firstRune) + s[1:]
}
func Substr(str string, start uint, length int) string {
	runes := []rune(str)
	strLen := len(runes)

	if start >= uint(strLen) {
		return ""
	}

	if length < 0 || int(start)+length > strLen {
		length = strLen - int(start)
	}

	return string(runes[start : start+uint(length)])
}

//Strrev 翻转字符串

func Strrev(str string) string {
	runes := []rune(str)
	runesLen := len(runes)

	for i := 0; i < runesLen/2; i++ {
		runes[i], runes[runesLen-1-i] = runes[runesLen-1-i], runes[i]
	}

	return string(runes)
}

//StrShuffle 随机打乱一个字符串

func StrShuffle(str string) string {
	runes := []rune(str)
	runesLen := len(runes)

	// 1.20以后，不需要此行代码
	rand.Seed(time.Now().UnixNano())

	for i := runesLen - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
