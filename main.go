package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bschlaman/b-utils/pkg/utils"
	"github.com/bxcodec/faker/v3"
)

const (
	serverName     string = "Data-Generator"
	port           string = ":8080"
	configPath     string = "config.json"
	maxDataObjects int    = 1000000
)

type dataSchema struct {
	Timestamp string `faker:"timestamp" json:"timestamp"`
	FirstName string `faker:"first_name" json:"fname"`
	LastName  string `faker:"last_name" json:"lname"`
	Email     string `faker:"email" json:"email"`
}

func genData() (interface{}, error) {
	ds := dataSchema{}
	if err := faker.FakeData(&ds); err != nil {
		return nil, err
	}
	return ds, nil
}

func genDataMany(n int) ([]interface{}, error) {
	data := make([]interface{}, n)
	for i := 0; i < n; i++ {
		d, err := genData()
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	return data, nil
}

// Handlers
func dataHandle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// default to 1 data object
		numDataObjects := 1
		parsed, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err == nil && parsed > 0 {
			numDataObjects = parsed
		}

		if numDataObjects > maxDataObjects {
			http.Error(w, fmt.Sprintf("can only request up to %d objects", maxDataObjects), http.StatusBadRequest)
			return
		}

		data, err := genDataMany(numDataObjects)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
		res, err := json.Marshal(&data)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(res))
	})
}

func initHandlers() {
	http.Handle("/data", dataHandle())
	http.Handle("/echo", utils.EchoHandle())
}

func main() {
	fmt.Println("starting", serverName, "on port", port)
	initHandlers()
	fmt.Println("initialized handlers")
	http.ListenAndServe(port, nil)
}
