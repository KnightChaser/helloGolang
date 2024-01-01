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

	data := make([]Article, 1) // Create a Article type slice with size of 1
	data[0].ArticleId = 80007741
	data[0].Author.UserId = "kn1ght"
	data[0].Author.UploadedTimestamp = time.Date(2024, 1, 1, 0, 13, 44, 665, time.UTC)
	data[0].Title = "Happy new year bros"
	data[0].Content = ":)"
	data[0].UpvoterIds = []string{"exus1a1", "pglogistics", "akz238", "rhd777", "phdor1research"}

	data[0].Comments = make([]ArticleComment, 2)
	data[0].Comments[0].UserId = "exus1a1"
	data[0].Comments[0].Content = "Happy apple pie!"
	data[0].Comments[0].UploadedTimestamp = time.Date(2024, 1, 1, 0, 14, 17, 444, time.UTC)
	data[0].Comments[1].UserId = "talul4h"
	data[0].Comments[1].Content = "Come to my house tomorrow evening"
	data[0].Comments[1].UploadedTimestamp = time.Date(2024, 1, 1, 0, 15, 03, 111, time.UTC)

	// Just for pretty print
	doc, _ := json.MarshalIndent(data, "", "	")
	fmt.Println(string(doc))

}
