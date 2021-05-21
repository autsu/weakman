package test

import (
	"log"
	"testing"
	dao2 "vote/v1/dao"
)

func init() {
	log.SetFlags(log.Llongfile | log.Ldate)
}

func TestAddCollege(t *testing.T) {
	n, err := dao2.CollegeAdd("test")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(n)
}

func TestQueryCollege(t *testing.T) {
	college, err := dao2.CollegeQueryByName("test1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", college)
}

func TestQueryCollegeByIdWithProfession(t *testing.T) {
	college, err := dao2.CollegeQueryByIdWithProfession("1")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", college)
	for _, p := range college.Professions {
		log.Println(p)
	}
}

func TestSoftDelete(t *testing.T) {
	n, err := dao2.CollegeSoftDelete("2")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(n)
}

func TestSoftDeleteMore(t *testing.T) {
	n, err := dao2.CollegeSoftDeleteMore([]string{"2", "3"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(n)
}

func TestQueryAllCollege(t *testing.T) {
	cs, err := dao2.CollegeQueryAll()
	if err != nil {
		log.Fatalln(err)
	}
	for _, c := range cs {
		log.Println(c)
	}
}
