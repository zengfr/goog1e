package goutil

import (
	"fmt"
	"strings"
)

// BitAutoString 将字节自动输出成合适的单位内容。
func BitAutoString(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%0.2fB", float64(size))
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%0.2fKB", float64(size)/(1024))
	}
	if size < 1024*1024*1024 {
		return fmt.Sprintf("%0.2fMB", float64(size)/(1024*1024))
	}
	if size < 1024*1024*1024*1024 {
		return fmt.Sprintf("%0.2fGB", float64(size)/(1024*1024*1024))
	}
	return fmt.Sprintf("%0.2fTB", float64(size)/(1024*1024*1024*1024))
}

var chineseNumber = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
var chineseCompany = []string{"个", "拾", "佰", "仟", "万"}

// NumberToChineseCapital 数字转成中文大写数字
func NumberToChineseCapital(number int64) string {
	var all []string
	for i := 0; true; i++ {
		temp := number % 10000
		number = number / 10000
		all = append(all, fmt.Sprint(numberChineseCapital(int(temp)), i))
		if number <= 0 {
			break
		}
	}
	var rt = ""
	for i := len(all) - 1; i >= 0; i-- {
		rt += all[i]
	}
	rt = strings.ReplaceAll(rt, "0", "")
	rt = strings.ReplaceAll(rt, "1", "万")
	rt = strings.ReplaceAll(rt, "2", "亿")
	rt = strings.ReplaceAll(rt, "3", "万")
	rt = strings.ReplaceAll(rt, "4", "兆")
	rt = strings.ReplaceAll(rt, "5", "万")
	rt = strings.ReplaceAll(rt, "6", "亿")
	if strings.HasPrefix(rt, "零") {
		rt = rt[len("零"):]
	}
	if strings.HasSuffix(rt, "零") {
		rt = rt[:len(rt)-len("零")]
	}
	rt = strings.ReplaceAll(rt, "零亿", "亿")
	rt = strings.ReplaceAll(rt, "零万", "万")
	rt = strings.ReplaceAll(rt, "亿万", "亿")
	if strings.HasPrefix(rt, "壹拾") {
		rt = rt[len("壹"):]
	}
	return rt
}

// numberChineseCapital 最多4位数字转大写，输出结果前面带零
func numberChineseCapital(number int) string {
	str := fmt.Sprint(number)
	var all = ""
	var j = 0
	for i := len(str) - 1; i >= 0; i-- {
		all = fmt.Sprintf("%v%v%v", chineseNumber[str[i]-'0'], chineseCompany[j], all)
		j++
	}
	for {
		temp := all
		all = strings.ReplaceAll(all, "零拾", "零")
		all = strings.ReplaceAll(all, "零佰", "零")
		all = strings.ReplaceAll(all, "零零", "零")
		if temp == all {
			break
		}
	}
	all = strings.ReplaceAll(all, "零个", "")
	all = strings.ReplaceAll(all, "个", "")
	if len(str) != 4 {
		return fmt.Sprint("零", all)
	}
	return all
}
