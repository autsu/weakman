package test

import (
	"log"
	"testing"
	"vote/v2/dao"
	"vote/v2/tool"
)

func TestUtoB(t *testing.T) {
	topic, _ := dao.TopicQueryById("7")
	utoB := tool.UtoB(topic.Deadline)
	log.Println(utoB)
}
