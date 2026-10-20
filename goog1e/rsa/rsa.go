package rsa

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var SysConfigDir = "./"

const (
	MD5_SALT = "cnhuashao"
)

var (
	publicKey  []byte
	privateKey []byte
)

// InitEncrypt
// @Description: 初始化秘钥
func InitEncrypt(pubKey, priKey string) {
	publicKey = []byte(pubKey)
	privateKey = []byte(priKey)

	//GenerateRSAKey(1024 * 4)
}

// 生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(SysConfigDir + "private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)
	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(SysConfigDir + "public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

func RSA_DecryptLong(cipherText []byte, paths ...string) ([]byte, error) {
	path := SysConfigDir + "private.pem"
	if len(paths) > 0 {
		path = paths[0]
	}
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	path2 := SysConfigDir + "public.pem"
	if len(paths) > 0 {
		path2 = paths[0]
	}
	//打开文件
	file2, err := os.Open(path2)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	//读取文件的内容
	info2, _ := file2.Stat()
	buf2 := make([]byte, info2.Size())
	file2.Read(buf2)
	//pem解码
	block2, _ := pem.Decode(buf2)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block2.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	pubKey := publicKeyInterface.(*rsa.PublicKey)

	partLen := pubKey.N.BitLen() / 8
	chunks := split(cipherText, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		//对密文进行解密
		plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(plainText)
	}
	//返回明文
	return []byte(buffer.String()), nil
}
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

// RSA_Encrypt
// @Description: RSA加密，当前端传过来时需要使用
// @param plainText
// @param path
// @return []byte
func RSA_Encrypt(plainText []byte, paths ...string) []byte {
	path := SysConfigDir + "public.pem"
	if len(paths) > 0 {
		path = paths[0]
	}
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	log.Println(string(buf))
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	pubKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

// RSA_Decrypt
// @Description: RSA解密
// @param cipherText
// @param paths
// @return []byte
// @return error
func RSA_Decrypt(cipherText []byte, paths ...string) ([]byte, error) {
	path := SysConfigDir + "private.pem"
	if len(paths) > 0 {
		path = paths[0]
	}
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, cipherText)
	if err != nil {
		return nil, err
	}
	//返回明文
	return plainText, nil
}

// Md5Crypt 加盐MD5加密
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// MD5Bytes 数据块MD5计算函数
func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return fmt.Sprintf("%x", ret)
}

// MD5 计算字符串MD5值
func MD5(s string) string {
	return MD5Bytes([]byte(s))
}

// EncryptDES_CBC CBC加密
func EncryptDES_CBC(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

// CBC解密
func DecryptDES_CBC(src, key string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

// ECB加密
func EncryptDES_ECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

// ECB解密
func DecryptDES_ECB(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

// 明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// RsaEncrypt Rsa加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	pubKey := publicKeyInterface.(*rsa.PublicKey)

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, origData)
	return encryptedBytes, err
}

// RsaDecrypt Rsa解密
func RsaDecrypt(cipherText []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//解密密文
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
