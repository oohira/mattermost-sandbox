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

func post(url string, data map[string]string) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(str))
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

func example1(url string) (string, error) {
	return post(url, map[string]string{
		"text": "Hello, this is some text\nThis is more text. :tada:",
	})
}

func example2(url string) (string, error) {
	return post(url, map[string]string{
		"channel":  "town-square",
		"username": "test-automation",
		"icon_url": "https://www.mattermost.org/wp-content/uploads/2016/04/icon.png",
		"text":     "#### Test results for July 27th, 2017\n@channel please review failed tests.\n\n| Component  | Tests Run   | Tests Failed                                   |\n|:-----------|:-----------:|:-----------------------------------------------|\n| Server     | 948         | :white_check_mark: 0                           |\n| Web Client | 123         | :warning: 2 [(see details)](http://linktologs) |\n| iOS Client | 78          | :warning: 3 [(see details)](http://linktologs) |",
	})
}

func main() {
	_, err := example1(WEBHOOK_URL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = example2(WEBHOOK_URL)
	if err != nil {
		log.Fatalln(err)
	}
}
