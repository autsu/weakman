package test

import (
	"log"
	"testing"
	"vote/tool"
)

func TestMd5(t *testing.T) {
	md5, err := tool.NewMD5("root")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(md5)
}

func TestMD5Equal(t *testing.T) {
	userInputPwd := "abcd12345"                 // 用户输入的原始数据
	dbPwd := "a99442d2a736365f5fe637e299b0e339" // 数据库中保存的加密数据
	ok := tool.MD5Equal(userInputPwd, dbPwd)
	log.Println(ok)
}
