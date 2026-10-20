package goutil

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// 找出字符串中所有数字, 如果没有匹配的，则返回nil
func FindNumsInString(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}

// CutStringByIndex 截取字符串，根据指定的开始下标和结束下标
// src:切片源
// start:开始下标，第一个为0
// end:结束位置，注意不包含结束的下标元素
func CutStringByIndex(src string, start int, end int) string {
	if start <= -1 {
		start = 0
	}
	if end <= start {
		return ""
	}
	length := len(src)
	if length <= start {
		return ""
	}
	if length >= end {
		return src[start:]
	}
	return src[start:end]
}

// CutSplitByIndex 截取切片，根据指定的开始下标和结束下标
// src:切片源
// start:开始下标，第一个为0
// end:结束位置，注意不包含结束的下标元素
func CutSplitByIndex[T IntUintBoolByteStringFloat](src []T, start int, end int) []T {
	if start <= -1 {
		start = 0
	}
	if end <= start {
		return []T{}
	}
	length := len(src)
	if length <= start {
		return []T{}
	}
	if length >= end {
		return src[start:]
	}
	return src[start:end]
}

// CutSplitByPage 截取切片，根据指定的页码和大小
// src:切片源
// page:页码，如果输入为0，则自动提升为1.
// size:页大小
func CutSplitByPage[T IntUintBoolByteStringFloat](src []T, page uint, size uint) []T {
	if page <= 0 {
		page = 1
	}
	start := (page - 1) * size //计算开始位置
	end := page * size         //计算结束位置
	if start > uint(len(src)) || end < start {
		return nil
	}
	if end > uint(len(src)) {
		return src[start:]
	}
	return src[start:end]
}

// CutStringBySubString 截取字符串，根据指定的开始子串和结束子串。如果不存在则返回空字符串。
// src:源字符串
// start:开始子串，如果 start 为空，则返回 end 前面那段。
// end:结束子串，如果 end 为空，则返回 start 后面那段。
// 如果 start 和 end 都为空，则返回原字符串 src 。
func CutStringBySubString(src string, start string, end string) string {
	if start == "" && end == "" {
		return src
	}
	if start == "" {
		before, _, _ := strings.Cut(src, end)
		return before
	}
	if end == "" {
		_, after, _ := strings.Cut(src, start)
		return after
	}
	_, after, found := strings.Cut(src, start)
	if found {
		before, _, f := strings.Cut(after, end)
		if f {
			return before
		}
	}
	return ""
}

// ToTitle 将字符串的首字母转换为大写
func ToTitle(str string) string {
	if len(str) <= 0 {
		return str
	}
	var _tempStr = []rune(strings.ToLower(str))
	if unicode.IsLetter(_tempStr[0]) {
		_tempStr[0] = unicode.ToUpper(_tempStr[0])
	}
	return string(_tempStr)
}

// ToUATitle 每个单词的首字母转换为大写
func ToUATitle(str string) string {
	if len(str) <= 0 {
		return str
	}
	var _tempStr = []rune(strings.ToLower(str))
	var first = true
	for i, d := range _tempStr {
		if unicode.IsLetter(d) {
			if first {
				_tempStr[i] = unicode.ToUpper(_tempStr[i])
				first = false
				continue
			}
			continue
		}
		if unicode.IsNumber(d) {
			if first {
				first = false
				continue
			}
			continue
		}
		first = true
	}
	return string(_tempStr)
}

// JoinString 拆分再合并字符串
// src：需要拆分的字符串模板，比如：4e6f9e19f397e8b3e42716aaf8f80589f1ed42c1272f7e7b25e5c5de7d88b2d1
// substr：需要往中间添加的子字符串，比如：-
// length：第 1 段的长度
// lengths：后续 n 段各自的长度
// 案例 JoinString(s, "-", 1, 2, 3, 4) 返回：4-e6-f9e-19f3
func JoinString(src string, substr string, length int, lengths ...int) string {
	lengths = append([]int{length}, lengths...)
	var rt strings.Builder
	for i, d := range lengths {
		if len(src) <= d {
			rt.WriteString(src)
			break
		}
		rt.WriteString(src[:d])
		if i != len(lengths)-1 {
			rt.WriteString(substr)
		}
		src = src[d:]
	}
	return rt.String()
}

// Join 拆分再合并
// src：需要拆分的模板，比如：[]rune("4e6f9e19f397e8b3e42716aaf8f80589f1ed42c1272f7e7b25e5c5de7d88b2d1")
// substr：需要往中间添加的子模板，比如：[]rune("-")
// length：第 1 段的长度
// lengths：后续 n 段各自的长度
// 案例 Join([]rune(s), []rune("-"), 1, 2, 3, 4) 返回：[]rune("4-e6-f9e-19f3")
func Join[T IntUintBoolByteStringFloat](src []T, substr []T, length int, lengths ...int) []T {
	lengths = append([]int{length}, lengths...)
	var rt []T
	for i, d := range lengths {
		if len(src) <= d {
			rt = append(rt, src...)
			break
		}
		rt = append(rt, src[:d]...)
		if i != len(lengths)-1 {
			rt = append(rt, substr...)
		}
		src = src[d:]
	}
	return rt
}

// PhoneKeepNumbersOnly 只保留下数字
func KeepNumbersOnly(phone string) string {
	var rt []byte
	for i := 0; i < len(phone); i++ {
		if phone[i] >= '0' && phone[i] <= '9' {
			rt = append(rt, phone[i])
		}
	}
	return string(rt)
}

// deleteAllSpace 删除掉空格，tab，回车，换行
func DeleteAllSpace(s string) string {
	var temp = make([]byte, len(s))
	var index = 0
	for i := 0; i < len(s); i++ {
		if s[i] < 0x80 {
			switch s[i] {
			case ' ', '\t', '\n', '\v', '\f', '\r':
				continue
			}
		}
		temp[index] = s[i]
		index++
	}
	s = strings.ReplaceAll(string(temp[:index]), string(0x85), "")
	s = strings.ReplaceAll(s, string(0xA0), "")
	return s
}

// 取字符串中间 如果未找到，原样返回。
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

// 取字符串中间，如果未找到返回空字符串
func BetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n < 0 {
		return ""

	}

	n += len(start)

	m := strings.Index(str[n:], end)

	if m < 0 {
		return ""
	}
	return str[n : n+m]
}

// 取字符串左侧
func LeftStr(str, posStr string, isContainsPosStr bool) string {
	n := strings.Index(str, posStr)
	if n < 0 {
		return ""
	}
	if isContainsPosStr {
		n += len(posStr)
	}

	// m := strings.Index(str[n:], end)

	// if m < 0 {
	// 	return ""
	// }
	return str[:n]
}

// 取字符串右侧
func RightStr(str, posStr string, isContainsPosStr bool) string {
	// 返回指定字符在此字符串中最后一次出现处的索引
	n := strings.LastIndex(str, posStr)
	if n < 0 {
		return ""
	}
	if !isContainsPosStr {
		n += len(posStr)
	}

	// m := strings.Index(str[n:], end)

	// if m < 0 {
	// 	return ""
	// }
	return str[n:]
}

// 补全网址协议头
func UrlPatchHttp(url string) (ret string) {
	if strings.Index(url, "http") < 0 && url != "" {
		var buffer bytes.Buffer
		buffer.WriteString("https:")
		buffer.WriteString(url)
		return buffer.String()
	}
	return url
}

// 截取字符串
func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}
func SubStrRange(s string, length int) string {
	var n, i int
	for i = range s {
		if n == length {
			break
		}

		n++
	}

	return s[:i]
}

// 截图字符串前N个字符内容
func SubStrRuneN(textStr string, n int) (retText string) {
	if len(textStr) <= 0 {
		return ""
	}
	retText = textStr
	textR := []rune(retText)
	if len(textR) > n {
		retText = string(textR[:n-1])
	}
	retText = strings.TrimSpace(retText)
	return retText
}
