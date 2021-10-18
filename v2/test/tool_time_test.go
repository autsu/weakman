package test

import (
	"log"
	"testing"
	"vote/v2/dao"
	"vote/v2/tool"
)

func TestUtoB(t *testing.T) {
	var dt *dao.TopicDao
	topic, _ := dt.QueryById("7")
	utoB := tool.UtoB(topic.Deadline)
	log.Println(utoB)
}
