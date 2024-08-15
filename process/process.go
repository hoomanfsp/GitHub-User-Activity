package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Start(user_name string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", user_name)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to make a request!", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("failed to get a successfull response!")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error, reading response body!")

	}
	var events []map[string]interface{}
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal("Failed to parse JSON", err)
	}

	// Print the parsed events
	for i, event := range events {
		fmt.Printf("Event %d:\n", i+1)
		fmt.Printf("  Type: %s\n", event["type"])
		fmt.Printf("  Repo: %s\n", event["repo"].(map[string]interface{})["name"])
		fmt.Println()
	}
}
