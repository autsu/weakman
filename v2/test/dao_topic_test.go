package test

import (
	"log"
	"testing"
	"vote/v2/dao"
)

func TestTopicQueryAllWithTopicSet(t *testing.T) {
	topics, err := dao.TopicQueryAllLimitWithTopicSet(2, 3)
	if err != nil {
		log.Fatalln(err)
	}
	for _, topic := range topics {
		log.Printf("%+v\n", topic)
	}
}

func TestTopicCount(t *testing.T) {
	count, err := dao.TopicCount()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(count)
}

func TestTopicQueryByTitle(t *testing.T) {
	topic, err := dao.TopicQueryByTitle("123")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", topic)
}