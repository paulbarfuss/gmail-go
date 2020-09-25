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
	for i := range msgs {
		fmt.Println(getmail.GetMessage(srv, user, msgs[i]))
	}
}
