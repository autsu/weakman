package test

import (
	"log"
	"testing"
	"vote/tool"
)

func TestConv(t *testing.T) {
	s := tool.Btos([]byte{'a', 'b', 'c'})
	log.Println(s)

	b := tool.Stob("123")
	log.Println(b)
}
