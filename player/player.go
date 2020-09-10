package player

import "net/http"

type player struct {
	playerName      string
	playerRealmname string
	client          *http.Client
}
