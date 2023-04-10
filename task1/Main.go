package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Response struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var (
	requests  = make(map[string]Request)
	responses = make(map[string]Response)
	mu        sync.Mutex
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Error parsing request JSON", http.StatusBadRequest)
		return
	}

	client := http.Client{}
	req, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error executing request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	response := Response{
		ID:      "requestId",
		Status:  resp.StatusCode,
		Headers: make(map[string]string),
		Length:  0,
	}

	for key, values := range resp.Header {
		response.Headers[key] = values[0]
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		response.Length = len(bodyBytes)
	}

	mu.Lock()
	requests["requestId"] = request
	responses["requestId"] = response
	mu.Unlock()

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
