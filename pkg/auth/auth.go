package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Retrieve a token, saves the token, then returns the generated client.
func CreateService() (*gmail.Service, error) {
	b, err := ioutil.ReadFile(path.Join(os.Getenv("HOME"), ".credentials.json"))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	token := getToken(config)

	tokenSource := config.TokenSource(context.Background(), token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))

	if err != nil {
		log.Fatalf("Unable to create Gmail service: %v", err)
	}
	return srv, nil

}

func getToken(config *oauth2.Config) *oauth2.Token {
	tokFile, err := os.Open(path.Join(os.Getenv("HOME"), ".token.json"))
	if err != nil {
		log.Fatalf("Unable to generate token: %v", err)
	}
	tok, err := tokenFromFile(tokFile.Name())
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile.Name(), tok)
	}
	return tok
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
