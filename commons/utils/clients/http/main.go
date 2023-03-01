package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type c struct{}

func (c) Get(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if res.Status == "200 OK" {
		type Comment struct {
			Id, Comment string
		}
		var comments Comment
		var text string
		var s = bufio.NewScanner(res.Body)
		for s.Scan() {
			//println(s.Text())
			text += s.Text()
		}
		err := json.Unmarshal([]byte(text), &comments)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Printf("%+v", comments)
		println()
		if err := s.Err(); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
func main() {
	client := &c{}
	client.Get("http://localhost:8080/api/comments")
}
