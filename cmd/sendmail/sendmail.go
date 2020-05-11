package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/paulbarfuss/gmail-go/pkg/auth"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func getMessage(service *gmail.Service, user string, msgId string) string {
	r, err := service.Users.Messages.Get(user, msgId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message: %v", err)
	}

	encoded := r.Payload.Body.Data
	decoded, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		log.Fatalf("Unable to decode message: %v", err)
	}
	messagePartBody := string(decoded)
	if len(messagePartBody) == 0 {
		return string("No message found in body")
	}

	return messagePartBody

}

// Creates a message for email

func main() {
	credentials := path.Join(os.Getenv("HOME"), ".credentials.json")
	b, err := ioutil.ReadFile(credentials)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := auth.GetClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"

	r, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}
	if len(r.Messages) == 0 {
		fmt.Println("No messages found.")
		return
	}
	fmt.Println("Messages:")
	for _, m := range r.Messages {

		message := getMessage(srv, user, m.Id)
		fmt.Println(message)
	}

}
