package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const SERVER_PORT = "8080"

const DEFAULT_URL string = "http://localhost:" + SERVER_PORT

func main() {
	go func() {
		var mux = http.NewServeMux() // for url patters  maps the corresponding registered handler
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			println("server:", request.Method)
			fmt.Fprintf(writer, `{"message":"Hello World"}`)

		})

		mux.HandleFunc("/id/", func(writer http.ResponseWriter, request *http.Request) {
			println("server:", request.Method)
			println("server :  query id :", request.URL.Query().Get("id"))
			println("server : content type :", request.Header.Get("Content-Type"))
			println("server :  headers :")
			for headerName, headerValue := range request.Header {
				fmt.Printf("\t %s = %s \n ", headerName,strings.Join(headerValue, ",")) 
			}

			reqBody, err := io.ReadAll(request.Body)
			if err != nil {
				println("server : error reading request body")
			}

			println("server : request body : ", string(reqBody))
                        fmt.Fprintf(writer, `{"message":"request processed"}`) 
                        time.Sleep(30*time.Second)

		})

		var server = http.Server{
			Addr:    ":" + SERVER_PORT,
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Println("server error:", err)
				os.Exit(1)
			}

		}
	}()

	time.Sleep(100 * time.Millisecond)

	jsonBody := []byte(`{"client_message":"Hello Server"}`)
	bodyReader := bytes.NewReader(jsonBody)
	request(http.MethodGet, DEFAULT_URL, nil)
	request(http.MethodPost, DEFAULT_URL+"/id/"+"?id=1234", bodyReader)

}

func request(method, requestURL string, bodyReader io.Reader) {

	request, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("error creating a request :{%s}\t\n", err)
		os.Exit(1)

	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Printf("error making  http request :{%s}\t\n", err)
		os.Exit(1)

	}

	println("Response : ", response.StatusCode)

	respBody, err := io.ReadAll(response.Body)

	if err != nil {
		println("client couldn't read the response body")
        
		os.Exit(1)
	}

	println("Response Body : ", string(respBody))

}
