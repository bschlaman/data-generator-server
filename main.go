package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bschlaman/b-utils/pkg/utils"
	"github.com/bxcodec/faker/v3"
)

const (
	serverName string = "Data-Generator"
	port       string = ":8080"
	configPath string = "config.json"
)

func genData() (interface{}, error) {
	data := struct {
		Timestamp string `faker:"timestamp" json:"timestamp"`
		FirstName string `faker:"first_name" json:"fname"`
		LastName  string `faker:"last_name" json:"lname"`
		Email     string `faker:"email" json:"email"`
	}{}
	if err := faker.FakeData(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func genDataMany(n int) ([]interface{}, error) {
	data := make([]interface{}, n)
	for i := 0; i < n; i++ {
		d, _ := genData()
		data[i] = d
	}
	return data, nil
}

// Handlers
func dataHandle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := genDataMany(2)
		res, _ := json.Marshal(&data)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(res))
	})
}

func initHandlers() {
	http.Handle("/data", dataHandle())
	http.Handle("/echo", utils.EchoHandle())
}

func main() {
	fmt.Println("starting", serverName)
	initHandlers()
	fmt.Println("initialized handlers")
	http.ListenAndServe(port, nil)
	fmt.Println("started", serverName, "on port", port)
}
