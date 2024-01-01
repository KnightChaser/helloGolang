package main

import (
	"encoding/json"
	"fmt"
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

	JSONDocument := `
	[{
		"ArticleId": 491803,
		"Author": {
		  "UserId": "kn1ght",
		  "UploadedTimestamp": "2024-01-01T12:34:56Z"
		},
		"Title": "Golang JSON >> C++ JSON",
		"Content": "I'm not kinda Golang person but handling JSON in Golang is better than C++ ngl",
		"UpvoterIds": ["oovp4817", "xhac3r", "fxsech"],
		"Comments": [
		  {
			"UserId": "oovp4817",
			"Content": "It truly is lmfao",
			"UploadedTimestamp": "2024-01-01T13:45:00Z"
		  },
		  {
			"UserId": "cpp17pablo",
			"Content": "You dare underrate the true power of C++?",
			"UploadedTimestamp": "2024-01-01T14:00:30Z"
		  }
		]
	  }]
	  `
	var data []Article
	json.Unmarshal([]byte(JSONDocument), &data)

	// Just for pretty print
	doc, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(string(doc))

	// Extract some data
	fmt.Println()
	fmt.Printf("data[0].ArticleId => %v(%T)\n", data[0].ArticleId, data[0].ArticleId)
	fmt.Printf("data[0].Author.UserId => %v(%T)\n", data[0].Author.UserId, data[0].Author.UserId)
	fmt.Printf("data[0].Comments[1].UploadedTimestamp => %v(%T)\n", data[0].Comments[1].UploadedTimestamp, data[0].Comments[1].UploadedTimestamp)
}
