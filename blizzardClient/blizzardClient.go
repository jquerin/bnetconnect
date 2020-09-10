package blizzardclient

import "net/http"

type bClient struct {
	clientID     string
	clientSecret string
	tokenURL     string
	client       *http.Client
}
