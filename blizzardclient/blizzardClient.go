package blizzardclient

import (
	ceq "bnetconnect/characterequipment"
	cp "bnetconnect/characterprofile"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// BnetClient represents blizzard client api
type BnetClient struct {
	Client             *http.Client
	CharacterProfile   cp.CharacterProfile
	CharacterEquipment ceq.CharacterEquipment
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
	profile := cp.CharacterProfile{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	b.CharacterProfile = profile
}

// GenerateCharacterEquipment method requests characterEquipment api
func (b *BnetClient) GenerateCharacterEquipment(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	equipment := ceq.CharacterEquipment{}
	err = json.NewDecoder(resp.Body).Decode(&equipment)
	b.CharacterEquipment = equipment
}
