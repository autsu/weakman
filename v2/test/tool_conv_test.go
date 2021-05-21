package test

import (
	"log"
	"testing"
	tool2 "vote/v1/tool"
)

func TestConv(t *testing.T) {
	s := tool2.Btos([]byte{'a', 'b', 'c'})
	log.Println(s)

	b := tool2.Stob("123")
	log.Println(b)
}
