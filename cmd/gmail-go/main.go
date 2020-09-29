package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/gmail/v1"

	"github.com/paulbarfuss/gmail-go/pkg/auth"
	"github.com/paulbarfuss/gmail-go/pkg/getmail"
)

func main() {
	srv, err := auth.CreateService()
	if err != nil {
		log.Fatalf("Could not create a new Gmail service: %v", err)
	}
	user := "me"
	messageIds := getmail.ListMessageIds(srv, user)

	jobs := make(chan string, len(messageIds))
	results := make(chan string, len(messageIds))

	for _, messageId := range messageIds {
		go worker(srv, user, messageId, jobs, results)
	}
	close(jobs)
	for j := 0; j < len(messageIds); j++ {
		fmt.Println(<-results)
	}
}

func worker(srv *gmail.Service, user string, messageId string, jobs <-chan string, results chan<- string) {
	results <- getmail.GetSnippet(srv, user, messageId)
}
