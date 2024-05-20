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

// $ Output:
// $ go run main.go
// 2024/05/20 07:48:18 Get "http://www.example.com": context deadline exceeded
// exit status 1

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

// $ Output:
// $ go run main.go
// 2024/05/20 07:47:13 Get "http://www.example.com": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
// exit status 1

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
	fmt.Println("Status", resp.Status, resp.Body)

	responsePerson := Person{}
	err = json.NewDecoder(resp.Body).Decode(&responsePerson)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resp Data:", responsePerson)
}

// sendJsonData :
// Sending JSON data is a common use case for POST requests. To send JSON data, we can use the http.Post method, but we need to make some changes:
//  1. We need to set the content type header to application/json - this tells the server that the request body is a JSON string
//  2. The request body needs to be a JSON string - we can use the JSON standard library to convert a Go struct to a JSON string
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

// postRequest :
// To make a POST request, we can use the http.Post method.
// This method takes in the URL, the content type, and the request body as parameters.
// The request body allows us to send data to the server, which is not possible with a GET request.
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

// getRequest()
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
	fmt.Println(resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error %d %s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(string(data))
	}
	fmt.Println(string(data))
}
