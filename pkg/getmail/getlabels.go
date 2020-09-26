package getmail

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/gmail/v1"
)

func GetLabels(srv *gmail.Service, user string) {

	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		log.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}
}
