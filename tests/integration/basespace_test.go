package tests

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ryanwentzel/go-basespace/basespace"
	"log"
	"net/url"
	"os"
)

func createClient() *basespace.Client {
	uri := os.Getenv("BASESPACE_API_URI")
	if uri == "" {
		log.Fatal("BaseSpace_API_URI not set!!!")
	}

	u, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Failed to parse uri: &v", uri)
	}

	token := os.Getenv("BASESPACE_AUTH_TOKEN")
	if token == "" {
		log.Fatal("BASESPACE_AUTH_TOKEN not set!!!")
	}

	transport := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	client := basespace.NewClient(transport.Client())
	client.ApiURL = u

	return client
}
