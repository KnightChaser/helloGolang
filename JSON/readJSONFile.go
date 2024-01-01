package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	type ArticleAuthor struct {
		UserId            string
		UploadedTimestamp time.Time
	}
	type ArticleComment struct {
		UserId            string
		Content           string
		UploadedTimestamp time.Time
	}
	type Article struct {
		ArticleId  uint64
		Author     ArticleAuthor
		Title      string
		Content    string
		UpvoterIds []string
		Comments   []ArticleComment
	}

	binaryJSONData, err := ioutil.ReadFile("JSON/postComment.json")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	// Read the JSON file and convert it to the Golang binary data
	var data []Article
	json.Unmarshal(binaryJSONData, &data)
	fmt.Println(data)

	// Extract specific JSON field
	fmt.Printf("The author of Article[0] post => %v\n", data[0].Author.UserId)
}
