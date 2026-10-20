package goutil

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type UtilService struct {
}

// uuid
func (this *UtilService) Uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:9], uuid[9:12], uuid[12:15]), nil
}

// 随机打乱字符串
func (this *UtilService) RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(t))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}

// 随机打乱切片
func (this *UtilService) RandomSlice2(stringss *[]interface{}) {

	for i := len(*stringss) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		(*stringss)[i], (*stringss)[num] = (*stringss)[num], (*stringss)[i]
	}

}

// 2个切片的交集
func (this *UtilService) SliceIntersect(s1 []string, s2 []string) []string {
	s := make([]string, 0)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				s = append(s, v1)
			}
		}
	}
	return s
}

// 2个切片的差集
func (this *UtilService) SliceDiff(s1 []string, s2 []string) []string {
	s := make([]string, 0)
	for _, v1 := range s1 {

		if !this.InArray(s2, v1) {
			s = append(s, v1)
		}

	}
	return s
}

// 判断切片中是否存在某个元素
func (this *UtilService) InArray(s1 []string, val string) bool {
	for _, v := range s1 {
		if v == val {
			return true
		}
	}
	return false
}
