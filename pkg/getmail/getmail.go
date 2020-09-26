// Package getmail checks gmail for messages
package getmail

import (
	"encoding/base64"

	log "github.com/sirupsen/logrus"

	"google.golang.org/api/gmail/v1"
)

func ListMessages(srv *gmail.Service, user string) []string {

	var msgs []string
	r, err := srv.Users.Messages.List(user).LabelIds("INBOX").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message list: %v", err)
	}

	for _, msg := range r.Messages {
		msgs = append(msgs, msg.Id)
	}
	return msgs

}

func GetMessage(srv *gmail.Service, user string, msgId string) string {
	r, err := srv.Users.Messages.Get(user, msgId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve message: %v", err)
	}

	encoded := r.Payload.Body.Data
	decoded, err := base64.URLEncoding.DecodeString(string(encoded))
	if err != nil {
		log.Fatalf("Unable to decode message: %v", err)
	}
	messagePartBody := string(decoded)
	if len(messagePartBody) == 0 {
		return string("No message found in body")
	}

	return messagePartBody
}
