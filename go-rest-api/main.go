package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	getRequest()
	postRequest()
	sendJsonData()
	parsingJsonData()
	setHeaders()
	setTimeouts()
	setSingleRequestTimeout()
}

// Person is a struct that represents the data we will send in the request body
type Person struct {
	Name string
	Age  int
}

func setSingleRequestTimeout() {
	url := "http://www.example.com"

	// create a new context instance with a timeout of 50 milliseconds
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)

	// create a new request using the context instance
	// now the context timeout will be applied to this request as well
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
}

func setTimeouts() {
	url := "http://www.example.com"

	// set a timeout of 50 milliseconds. This means that if the server does not
	// respond within 50 milliseconds, the request will fail with a timeout error
	http.DefaultClient.Timeout = 50 * time.Millisecond
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}

func setHeaders() {
	url := "http://localhost:3000"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// add a custom header to the request
	// here we specify the header name and value as arguments
	req.Header.Add("X-Custom-Header", "custom-value")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print all the response headers
	fmt.Println("Response headers:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func parsingJsonData() {
	url := "http://localhost:3000"
	p := Person{
		Name: "Roushan Narayan Sah",
		Age:  25,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status", resp.Status)

	responsePerson := Person{}
	err = json.NewDecoder(resp.Body).Decode(&responsePerson)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resp Data:", responsePerson)
}

func sendJsonData() {
	url := "http://localhost:3000"
	p := Person{
		Name: "Roushan Narayan Sah",
		Age:  25,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Status", resp.Status)
}

func postRequest() {
	url := "http://localhost:3000"

	body := strings.NewReader("This is the request body")
	resp, err := http.Post(url, "text/plain", body)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
}

// Making a GET Request
// For certain common request methods, like GET, and POST, the http package provides helper methods that
// make it easier to make requests.
func getRequest() {
	url := "http://www.example.com"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error %d %s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(string(data))
	}
}
