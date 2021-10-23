package test

import (
	"log"
	"testing"
	"vote/dao"
	"vote/tool"
)

func TestUtoB(t *testing.T) {
	var dt *dao.TopicDao
	topic, _ := dt.QueryById("7")
	utoB := tool.UtoB(topic.Deadline)
	log.Println(utoB)
}
