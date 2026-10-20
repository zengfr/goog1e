package goutil

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// hash plaintext with SHA-1
func EncryptSha1(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// hash plaintext with md5
func EncryptMd5(source string) (cryptext string) {
	if len(source) <= 0 {
		return ""
	}
	// 使用MD5加密
	signBytes := md5.Sum([]byte(source))
	// 把二进制转化为大写的十六进制
	cryptext = hex.EncodeToString(signBytes[:])
	return
}

// hash plaintext with md5
func EncryptMd5Byte(source []byte) (cryptext string) {
	// 使用MD5加密
	signBytes := md5.Sum(source)
	// 把二进制转化为大写的十六进制
	cryptext = hex.EncodeToString(signBytes[:])
	return
}

// https://www.jianshu.com/p/02da10ca45d1
// hmac-md5加密
func EncryptHmacmd5(src, key string) (cryptext string) {
	// m:=md5.New()
	m := hmac.New(md5.New, []byte(key))
	m.Write([]byte(src))
	cryptext = hex.EncodeToString(m.Sum(nil))
	return
}
func EncryptHmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func Base64StdEncodingEncodeToString(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}
func Base64StdEncodingDecode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
func Gzip(data []byte) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	//fmt.Println("gzip size:", len(b.Bytes()))
}

func UnGzip(byte []byte) []byte {
	b := bytes.NewBuffer(byte)
	r, _ := gzip.NewReader(b)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	//fmt.Println("ungzip size:", len(undatas))
	return undatas
}

func JsonDecodeString(String string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(String), &jsonMap)
	return jsonMap
}

func JsonDecodeByte(bytes []byte) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	json.Unmarshal(bytes, &jsonMap)
	return jsonMap
}
func JsonEncodeMapToByte(stringMap map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(stringMap)
	if err != nil {
		return nil
	}
	return jsonBytes
}

/**
 * @note
 * 获取md5 hash串
 * @param string text 源串
 *
 * @return string
 */
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func JsonEncode(data interface{}) string {
	s, e := json.Marshal(data)
	if e != nil {
		return ""
	}
	return string(s)
}

func JsonDecode(data string, inter interface{}) error {
	return json.Unmarshal([]byte(data), inter)
}

func StructToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(data)
	_ = json.Unmarshal(j, &m)
	return m
}
