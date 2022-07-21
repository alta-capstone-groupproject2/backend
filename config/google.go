package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

var (
	googleOAuthConfig = &oauth2.Config{
		Endpoint: oauth2.Endpoint{AuthURL: "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: oauth2.AuthStyleInParams},
	}
)

func CalendarService() *calendar.Service {

	client := getClient(googleOAuthConfig)
	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}
	return srv
	// InsertEvent(srv)
	//UpdateEvent(srv, "qu112soiplifpvmuv0bt856oig")

}
func getClient(config *oauth2.Config) *http.Client {

	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config.Client(context.Background(), tok)
}

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
