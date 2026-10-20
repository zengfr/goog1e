package goutil

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"goog1e/hexdump"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func PrintStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
func ArraySplitByValue(data []byte, value byte) [][]byte {
	var result [][]byte
	for {
		index := bytes.IndexByte(data, value)
		if index == -1 {
			result = append(result, data)
			break
		} else {
			result = append(result, data[:index])
			data = data[index+1:]
		}
	}
	return result
}
func ArraySplitByLen(data []byte, splitLen int) [][]byte {
	var result [][]byte
	for i := 0; i < len(data); i += splitLen {
		end := i + splitLen
		if end > len(data) {
			end = len(data)
		}
		result = append(result, data[i:end])
	}
	return result
}

// appendSlices 将多个切片连接成一个新切片。
// 参数 slices 是要连接的切片列表，切片元素可以任意。
// 返回一个包含所有切片元素的新切片
func ArrayJoin[T any](slices ...[]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]T, totalLen)
	begin := 0
	for _, s := range slices {
		begin += copy(result[begin:begin+len(s)], s)
	}
	return result
}
func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
func BytesToInt32(bys []byte) int32 {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int32(data)
}

// bytes to int 16
func BytesTo16Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint16
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}
func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

// int to 4 bytes
func IntTo4Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint32(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// int to 2 bytes
func IntTo2Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint16(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}
func Println(a ...string) (n int, err error) {

	logrus.Println(Colored(a...))
	return 0, nil
}
func PrintHexStr(s string, dumper *hexdump.Dumper, tit string) {
	PrintHexBytes(StringToBytes(s), dumper, tit)
}
func PrintHexBytes(bs []byte, dumper *hexdump.Dumper, tit string) {

	if dumper == nil {
		dumper = hexdump.New()
	}
	dumper.Reset()
	l := len(bs)
	logrus.Println(fmt.Sprintf("%03d %03x %s %s", l, l, tit, strings.Repeat("-", 0x20*4+0x010-9-0x10-len(tit))))
	hex := dumper.Dump(bs)
	logrus.Println(Colored(hex))
}

// 生成随机UUID(RFC 4122)
// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (string, error) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		// log.Errorf(err, "Cannot generate UUID:%s")
		return "", err
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid, err
}

// 拼接转换 string 到 []byte
func StringToBytes(s string) []byte {
	var b bytes.Buffer
	b.WriteString(s)
	return b.Bytes()
}

// 拼接转换 []byte 到 string
func BytesToString(b []byte) string {
	var s strings.Builder
	s.Write(b)
	return s.String()
}

func B64encode(input string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(input)))
}

func B64decode(input string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(input)
	return string(decodedString)
}

func UnzipBytes(zippedBytes []byte) ([]byte, error) {
	reader := bytes.NewReader(zippedBytes)
	zipReader, err := zip.NewReader(reader, int64(len(zippedBytes)))
	if err != nil {
		return nil, err
	}
	if len(zipReader.File) == 0 {
		return nil, nil // No file to open / extract
	}
	f, err := zipReader.File[0].Open()
	if err != nil {
		panic(err)
	}
	p, err := io.ReadAll(io.Reader(f))
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 进行zlib压缩
func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// 进行zlib解压缩
func ZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

/*
*功能：加密
*描述：
*此函数接受一个字符串和一个密码密钥，并使用 AES 加密消息
*
*参数：
*byte[] key : 包含密钥的字节切片
*string message : 包含要加密的消息的字符串
*
*返回：
*字符串编码：包含编码用户输入的字符串
*错误错误：错误信息
 */
func AesEncrypt(key []byte, message string) (encoded string, err error) {
	//从输入字符串创建字节切片
	plainText := []byte(message)

	//使用密钥创建新的 AES 密码
	block, err := aes.NewCipher(key)

	//如果 NewCipher 失败，退出：
	if err != nil {
		return
	}

	// ^ 使密文成为大小为 BlockSize + 消息长度的字节切片,这样传值后修改不会更改底层数组
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	// ^ iv 是初始化向量 (16字节)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	// ^ 加密数据,给定加密算法用的密钥,以及初始化向量
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//返回以base64编码的字符串
	//return base64.RawStdEncoding.EncodeToString(cipherText), err //zfr
	return string(cipherText), err
}

/*
*功能：解密
*描述：
*此函数接受一个字符串和一个密钥，并使用 AES 将字符串解密为纯文本
*
*参数：
*[]byte key : 包含密钥的字节切片
*string secure : 包含加密消息的字符串
*
*返回：
*string decoded : 包含解密后的字符串
*错误错误：错误信息
 */
func AesDecrypt(key []byte, secure string) (decoded string, err error) {
	//删除 base64 编码：
	//cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	cipherText := StringToBytes(secure)
	//如果解码字符串失败，退出：
	if err != nil {
		return
	}

	//使用密钥和加密消息创建新的 AES 密码
	block, err := aes.NewCipher(key)

	//如果 NewCipher 失败，退出：
	if err != nil {
		return
	}

	//如果密文的长度小于 16 字节
	if len(cipherText) < aes.BlockSize {
		err = errors.New("text len too small")
		return
	}

	// ^ iv 是初始化向量 (16字节)
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//解密消息
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

func GetFileText(f string) string {
	var str = ""
	file, err := os.Open(f)
	if err != nil {
		logrus.Println("err = ", err)
	}
	//及时关闭 file 句柄，否则会有内存泄漏
	defer file.Close()
	//创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str2, err := reader.ReadString('\n') //读到一个换行就结束
		str += str2
		//logrus.Print(str)
		if err == io.EOF { //io.EOF 表示文件的末尾
			break
		}

	}
	logrus.Println("...")
	return str
}

func StrToNum(text string) uint32 {
	value, err := strconv.ParseUint(text, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(value)
}

func StrToNum64(text string) uint64 {
	value, err := strconv.ParseUint(text, 10, 32)
	if err != nil {
		panic(err)
	}
	return value
}

func UIntToBytes(x uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}
