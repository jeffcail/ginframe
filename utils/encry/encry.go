package encry

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
)

// Md5 md5加密
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

// Sha256 加密
func Sha256(str string) string {
	srcByte := []byte(str)
	hash := sha256.New()
	hash.Write(srcByte)
	hashBytes := hash.Sum(nil)
	sha256String := hex.EncodeToString(hashBytes)
	return sha256String
}

// Sha512 加密
func Sha512(str string) string {
	srcByte := []byte(str)
	hash := sha512.New()
	hash.Write(srcByte)
	hashBytes := hash.Sum(nil)
	sha256String := hex.EncodeToString(hashBytes)
	return sha256String
}

// FileMd5 文件加密
func FileMd5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}
	md5Str := hex.EncodeToString(hash.Sum(nil))
	return md5Str, nil
}

// ScryptPasswd 密码 + 盐（一串随机数） 再Hash加密
func ScryptPasswd(password string) string {
	const cost = 10
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashPwd)
}

// ComparePassword 密码比较
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
