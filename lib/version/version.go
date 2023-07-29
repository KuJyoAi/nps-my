package version

import (
	"ehang.io/nps/lib/crypt"
)

const VERSION = "0.26.10"

var (
	CONN_TEST_NEW = crypt.GetRandomString(32)
	Server_Random = crypt.GetRandomString(8)
)

// Compulsory minimum version, Minimum downward compatibility to this version
func GetVersion() string {
	return "0.26.0"
}

/*
做一个基础的MD5加密
*/
func Base_Md5() (a string) {
	s := GetVersion()
	a = crypt.Md5(s + CONN_TEST_NEW)
	return
}

func Aes_Encode(orige string, key []byte) (NewData []byte, err error) {
	return crypt.AesEncrypt([]byte(orige), []byte(key))
}
func Aes_Decode(NewData string, key []byte) (orige []byte, err error) {
	return crypt.AesDecrypt([]byte(NewData), key)
}
