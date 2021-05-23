package test

import (
	"log"
	"testing"
	"vote/v2/dao"
)

func TestTopicOptionQueryById(t *testing.T) {
	option, err := dao.TopicOptionQueryById("1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(option)
}