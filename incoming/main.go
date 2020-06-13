package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	WEBHOOK_URL = "http://localhost:8065/hooks/ftcxcu4aypfutqby89yhpe7kze"
)

func hello(url string) (string, error) {
	data, err := json.Marshal(map[string]string{
		"text": "Hello, this is some text\nThis is more text. :tada:",
	})
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Unexpected status code: %d\n", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	body, err := hello(WEBHOOK_URL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
