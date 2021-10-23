package test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"
	"vote/model"
)

func TestDecodeJsonToBJT(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type T struct {
			Deadline model.BJT
		}
		body := r.Body
		js, _ := io.ReadAll(body)
		//var b model.BJT
		var ttt T
		if err := json.Unmarshal(js, &ttt.Deadline); err != nil {
			log.Fatalln(err)
		}
		log.Println(ttt)
		//log.Println(b)
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalln(err)
	}
}

func Test111(t *testing.T) {
	type T struct {
		Deadline model.BJT
	}
	//body := r.Body
	//js, _ := io.ReadAll(body)
	//var b model.BJT
	var ttt T
	if err := json.Unmarshal([]byte("2006-01-02 15:04:05"), &ttt.Deadline); err != nil {
		log.Fatalln(err)
	}
	log.Println(ttt)
}
