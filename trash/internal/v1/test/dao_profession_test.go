package test

import (
	"log"
	"testing"
	dao2 "vote/trash/internal/v1/dao"
)

func TestProfessionQueryByIdWithGrade(t *testing.T) {
	p, err := dao2.ProfessionQueryByIdWithGrade("1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(p)
	for _, g := range p.Grades {
		log.Println(g)
	}
}