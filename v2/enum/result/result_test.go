package result

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestResultWithResponse(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := New(SUCCESS, SUCCESS.String(), nil)
		jr, err := json.Marshal(res)
		if err != nil {
			io.Copy(w, strings.NewReader(err.Error()))
		}
		w.Write(jr)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
		return
	}
}
