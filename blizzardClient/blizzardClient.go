package blizzardclient

import (
	cp "bnetconnect/characterProfile"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// BnetClient represents blizzard client api
type BnetClient struct {
	Client           *http.Client
	CharacterProfile cp.Characterprofile
}

// CreateClient creates http client to send out requests
func CreateClient(clientID string, clientSecret string, tokenURL string) *http.Client {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	client := conf.Client(context.Background())
	return client
}

// GenerateCharacterProfile method requests characterProfile api
func (b *BnetClient) GenerateCharacterProfile(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	profile := cp.Characterprofile{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	b.CharacterProfile = profile
}
