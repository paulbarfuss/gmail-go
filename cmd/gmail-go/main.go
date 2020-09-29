package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/paulbarfuss/gmail-go/pkg/auth"
	"github.com/paulbarfuss/gmail-go/pkg/getmail"
)

func main() {
	srv, err := auth.CreateService()
	if err != nil {
		log.Fatalf("Could not create a new Gmail service: %v", err)
	}
	user := "me"
	msgs := getmail.ListMessages(srv, user)
	snippets := make(chan string)
	go func(msgs []string) {
		for _, snippet := range msgs {
			unpack, _ := getmail.GetSnippet(srv, user, snippet)
			snippets <- unpack
		}
	}(msgs)
	fmt.Println(snippets)
}
