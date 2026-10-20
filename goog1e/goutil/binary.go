package goutil

import "fmt"

const (
	BitTrue0 = 0x80 //1000  0000
	BitTrue1 = 0x40 //0100 0000
	BitTrue2 = 0x20 //0010 0000
	BitTrue3 = 0x10 //0001 0000
	BitTrue4 = 0x08 //0000 1000
	BitTrue5 = 0x04 //0000 0100
	BitTrue6 = 0x02 //0000 0010
	BitTrue7 = 0x01 //0000 0001

	BitFalse0 = 0x7F //0111 1111
	BitFalse1 = 0xBF //1011 1111
	BitFalse2 = 0xDF //1101 1111
	BitFalse3 = 0xEF //1110 1111
	BitFalse4 = 0xF7 //1111 0111
	BitFalse5 = 0xFB //1111 1011
	BitFalse6 = 0xFD //1111 1101
	BitFalse7 = 0xFE //1111 1110
)

// ByteToBinaryString 将byte数组，转换为二进制字符串
func ByteToBinaryString(all []byte) string {
	var result string
	for _, d := range all {
		result += fmt.Sprintf("%08b", d)
	}
	return result
}

// BitIndexStatus 判断byte数组的第多少个二进制位是1还是0，1返回true，0返回false。
// index 下标，0代表第一个；1代表第二个...。
func BitIndexStatus(all []byte, index uint) bool {
	return BitIndex(all[index/8], index%8)
}

// BitIndex 返回 c 的二进制第多少位为1还是0，1返回true，0返回false
// index 下标，0代表第一个；1代表第二个...。
func BitIndex(c uint8, index uint) bool {
	index++
	if index > 8 {
		return false
	}
	switch index % 8 {
	case 1:
		return c&BitTrue0 == BitTrue0
	case 2:
		return c&BitTrue1 == BitTrue1
	case 3:
		return c&BitTrue2 == BitTrue2
	case 4:
		return c&BitTrue3 == BitTrue3
	case 5:
		return c&BitTrue4 == BitTrue4
	case 6:
		return c&BitTrue5 == BitTrue5
	case 7:
		return c&BitTrue6 == BitTrue6
	default:
		return c&BitTrue7 == BitTrue7
	}
}
