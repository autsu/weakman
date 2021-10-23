package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"vote/errno"
)

func NewMD5(s string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		logrus.Warnf("%s: %s\n", errno.EncryptPasswordError, err)
		return "", errno.EncryptPasswordError
	}

	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

// MD5Equal  row 加密后是否与 md 匹配
func MD5Equal(row, md string) bool {
	m, _ := NewMD5(row)
	return m == md
}
