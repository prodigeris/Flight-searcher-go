package main

import (
	"encoding/json"
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	"io"
	"log"
	"net/http"
)

type Inquiry struct {
	WeekendCount int `json:"weekend_count"`
}

func main() {

	_, mqch, err := common.GetRabbitClient()
	if err != nil {
		log.Fatalf("Rabbit cannot be started")
	}
	common.DeclareQueue(mqch, common.InquiriesQueue)

	http.HandleFunc("/inquiry", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var inquiry Inquiry

		err := json.NewDecoder(r.Body).Decode(&inquiry)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(r.Body)

		err = publishInquiry(mqch, inquiry)
		if err != nil {
			http.Error(w, "Message cannot be published", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		responseJSON := `{"success": true}`
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(responseJSON))
		if err != nil {
			return
		}
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
