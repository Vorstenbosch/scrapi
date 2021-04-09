package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var server http.Server

func TestHealthEndpoint(t *testing.T) {
	// Given
	expectedResult := "I am healthy :)"

	// When
	// Server is already running via the setup method

	// Then
	response, err := http.Get("http://localhost:8000/v1/healthy")

	if err != nil {
		t.Errorf("Getting the 'healthy' endpoint resulted in error '%v'", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if string(body) != expectedResult {
		t.Errorf("The 'healthy' endpoint did not return the expected result '%s' but returned '%s'", expectedResult, body)
	}
}

func TestScrapeResultEndpoint(t *testing.T) {
	// Given
	expectedResult := "\"Error\":null"

	// When
	// Server is already running via the setup method
	time.Sleep(10 * time.Second)

	// Then
	response, err := http.Get("http://localhost:8000/v1/scrape-result")

	if err != nil {
		t.Errorf("Getting the 'scrape-result' endpoint resulted in error '%v'", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if !strings.Contains(string(body), expectedResult) {
		t.Errorf("The 'scrape-result' endpoint did not return a json object containing '%s' but was '%v'", expectedResult, body)
	}
}

func setup() {
	server := &http.Server{Addr: ":5555"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<div>Hello world</div>")
	})

	go func() {
		// always returns error. ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	yamlConfig := `
scrapeintervalinseconds: 5
scrapeendpoints: 
  - endpoint: http://localhost:5555
    selectors:
      - name: test
        typeofselector: xpath
        value: //div
`

	go startServer([]byte(yamlConfig))
}

func tearDown() {
	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	tearDown()
	os.Exit(exitCode)
}
