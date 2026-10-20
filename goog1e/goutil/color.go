package goutil

import (
	"fmt"
	"math/rand"
	"strings"
)

var Reset = "\033[0m"

/////////////
// Special //
/////////////

var Bold = "\033[1m"
var Underline = "\033[4m"

var Black = "\033[30m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

///////////////////////
// Background colors //
///////////////////////

var BlackBackground = "\033[40m"
var RedBackground = "\033[41m"
var GreenBackground = "\033[42m"
var YellowBackground = "\033[43m"
var BlueBackground = "\033[44m"
var PurpleBackground = "\033[45m"
var CyanBackground = "\033[46m"
var GrayBackground = "\033[47m"
var WhiteBackground = "\033[107m"
var Colors = []string{Red, Green, Yellow, Blue, Purple, Cyan, Gray, White}

func colored(c string, arr ...string) string {

	return c + strings.Join(arr, ", ") + Reset
}
func Colored(arr ...string) string {
	l := len(Colors)
	c := rand.Intn(l * 100)
	return Colors[c%l] + strings.Join(arr, ", ") + Reset
}
func PrintHexdump(data []byte) {
	fmt.Println(Hexdump(0, data))
}
func Hexdump(offset int, data []byte) string {
	var result string
	red := Red

	for len(data) > 0 {
		l := len(data)
		if l > 32 {
			l = 32
		}
		work := data[:l]
		data = data[l:]

		var workHex string
		var workAscii string
		for i := 0; i < 32; i++ {
			m := byte(0)
			valid := i < len(work)
			delta := false
			if valid {
				m = work[i]

			}

			if valid {
				if m < 32 || m > 126 {
					m = '.'
					delta = true
				}
				if delta {
					workHex += colored(red, fmt.Sprintf("%02x ", m))
				} else {
					workHex += fmt.Sprintf("%02x ", m)
				}

				if delta {
					workAscii += colored(fmt.Sprintf("%c", m))
				} else {
					workAscii += fmt.Sprintf("%c", m)
				}
			} else {
				workHex += "   "
				workAscii += " "
			}
			if i%8 == 7 {
				workHex += " "
			}
		}

		result += fmt.Sprintf("%08x  %s|%s|\n", offset, workHex, workAscii)
		offset += l
	}

	return result
}
