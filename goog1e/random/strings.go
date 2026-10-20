package random

import (
	"fmt"
)

type TemplateString string

const (
	TemplateString_0_9         TemplateString = `1234567890`
	TemplateString_a_z         TemplateString = `abcdefghijklmnopqrstuvwxyz`
	TemplateString_A_Z         TemplateString = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_0_9_a_f     TemplateString = `1234567890abcdef`
	TemplateString_0_9_a_z     TemplateString = `1234567890abcdefghijklmnopqrstuvwxyz`
	TemplateString_0_9_A_F     TemplateString = `1234567890ABCDEF`
	TemplateString_0_9_A_Z     TemplateString = `1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_a_z_A_Z     TemplateString = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	TemplateString_0_9_a_z_A_Z TemplateString = `1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

// String 生成随机的字符串
func String(t TemplateString, length int) string {
	if length <= 0 {
		return ""
	}
	var rt string
	var tLen = len(t)
	for i := 0; i < length; i++ {
		rt += fmt.Sprintf("%c", t[Rand.Intn(tLen)])
	}
	return rt
}

// Password 生成包含大写，小写，数字的字符串，以及指定特殊字符的密码
func Password(length int, otherChar string) string {
	if length <= 0 {
		return ""
	}
	var rt string
	if len(otherChar) > 0 {
		rt += fmt.Sprintf("%c", otherChar[Rand.Intn(len(otherChar))])
	}
	rt += String(TemplateString_0_9, 1)
	rt += String(TemplateString_a_z, 1)
	rt += String(TemplateString_A_Z, 1)
	if len(otherChar) > 0 {
		rt += String(TemplateString_0_9_a_z_A_Z+TemplateString(otherChar), length-4)
	} else {
		rt += String(TemplateString_0_9_a_z_A_Z, length-4)
	}
	temp := []byte(rt)
	temp = Shuffle(temp)
	return string(temp)
}
