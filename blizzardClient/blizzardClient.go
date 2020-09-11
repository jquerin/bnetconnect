package blizzardclient

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

type bClient struct {
	clientID     string
	clientSecret string
	tokenURL     string
	client       *http.Client
}

func CreateClient(clientID string, clientSecret string, tokenURL string) *http.Client {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	client := conf.Client(context.Background())
	return client
}
