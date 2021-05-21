package tool

import "unsafe"

func Stob(s string) []byte {
	b := *(*[]byte)(unsafe.Pointer(&s))
	return b
}

func Btos(b []byte) string {
	s := *(*string)(unsafe.Pointer(&b))
	return s
}
