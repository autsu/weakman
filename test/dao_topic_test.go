package test

import (
	"log"
	"testing"
	"vote/dao"
)

func TestTopicQueryAllWithTopicSet(t *testing.T) {
	var dt *dao.TopicDao
	topics, err := dt.QueryAllLimitWithTopicSet(2, 3)
	if err != nil {
		log.Fatalln(err)
	}
	for _, topic := range topics {
		log.Printf("%+v\n", topic)
	}
}

func TestTopicCount(t *testing.T) {
	var dt *dao.TopicDao
	count, err := dt.Count()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(count)
}

func TestTopicQueryByTitle(t *testing.T) {
	var dt *dao.TopicDao
	topic, err := dt.QueryByTitle("123")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", topic)
}

func TestTopicShowResultById(t *testing.T) {
	var dt *dao.TopicDao
	result, err := dt.ShowResultById("7")
	if err != nil {
		log.Fatalln(err)
	}
	for _, vo := range result {
		log.Printf("%+v\n", vo)
	}
}
