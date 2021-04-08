package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vorstenbosch/scrapy/scrapyboss"
)

func main() {
	// Getting the scrape config file
	configFile := os.Args[1]
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	startServer(data)
}

func startServer(data []byte) {
	config, err := scrapyboss.ParseConfig(data)
	if err != nil {
		log.Fatal(err)
	}

	sb := scrapyboss.NewScrapyBoss(config)
	sb.Start()

	r := mux.NewRouter()

	v1Router := r.PathPrefix("/v1").Subrouter()

	v1Router.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "I am happy :)")
	})

	v1Router.HandleFunc("/scrape-result", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(sb.GetScrapeData())
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		} else {
			http.Error(w, "my own error message", http.StatusForbidden)
		}
	})

	if os.Getenv("USE_TLS") == "true" {
		err := http.ListenAndServeTLS(":8000", os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY"), r)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = http.ListenAndServe(":8000", r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
