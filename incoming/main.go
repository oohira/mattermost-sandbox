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

func post(url string, data map[string]interface{}) (string, error) {
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
	return post(url, map[string]interface{}{
		"text": "Hello, this is some text\nThis is more text. :tada:",
	})
}

func example2(url string) (string, error) {
	return post(url, map[string]interface{}{
		"channel":  "town-square",
		"username": "test-automation",
		"icon_url": "https://www.mattermost.org/wp-content/uploads/2016/04/icon.png",
		"text": `#### Test results for July 27th, 2017
@channel please review failed tests.

| Component  | Tests Run   | Tests Failed                                   |
|:-----------|:-----------:|:-----------------------------------------------|
| Server     | 948         | :white_check_mark: 0                           |
| Web Client | 123         | :warning: 2 [(see details)](http://linktologs) |
| iOS Client | 78          | :warning: 3 [(see details)](http://linktologs) |`,
	})
}

func example3(url string) (string, error) {
	return post(url, map[string]interface{}{
		"channel":    "town-square",
		"username":   "Winning-bot",
		"icon_emoji": "+1",
		"text":       "#### We won a new deal!",
		"props": map[string]string{
			"card": "Salesforce Opportunity Information:\n\n [Opportunity Name](http://salesforce.com/OPPORTUNITY_ID)\n\n-Salesperson: **Bob McKnight** \n\n Amount: **$300,020.00**",
		},
	})
}

func main() {
	for _, f := range []func(string) (string, error){
		example1,
		example2,
		example3,
	} {
		body, err := f(WEBHOOK_URL)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(body)
	}
}
