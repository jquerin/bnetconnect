package blizzardclient

import (
	"context"
	"encoding/json"
	"fmt"
	ceq "github.com/mrgreenturtle/bnetconnect/characterequipment"
	cm "github.com/mrgreenturtle/bnetconnect/charactermedia"
	cp "github.com/mrgreenturtle/bnetconnect/characterprofile"
	pvp "github.com/mrgreenturtle/bnetconnect/characterpvp"
	pvps "github.com/mrgreenturtle/bnetconnect/characterpvpsummary"
	cs "github.com/mrgreenturtle/bnetconnect/characterspecialization"
	cstats "github.com/mrgreenturtle/bnetconnect/characterstats"
	ctits "github.com/mrgreenturtle/bnetconnect/charactertitles"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// BnetClient represents blizzard client api
type BnetClient struct {
	Client *http.Client
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
func (b *BnetClient) GenerateCharacterProfile(name string, realm string) cp.CharacterProfile {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	profile := cp.CharacterProfile{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	return profile
}

// GenerateCharacterTitles method requests characterTitles api
func (b *BnetClient) GenerateCharacterTitles(name string, realm string) ctits.CharacterTitles {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/titles?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	titles := ctits.CharacterTitles{}
	err = json.NewDecoder(resp.Body).Decode(&titles)
	if err != nil {
		fmt.Println(err)
	}
	return titles
}

// GenerateCharacterMedia method requests characterMedia api
func (b *BnetClient) GenerateCharacterMedia(name string, realm string) cm.CharacterMedia {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/character-media?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	media := cm.CharacterMedia{}
	err = json.NewDecoder(resp.Body).Decode(&media)
	return media
}

// GenerateCharacterSpecialization method requests characterSpecialization api
func (b *BnetClient) GenerateCharacterSpecialization(name string, realm string) cs.CharacterSpecialization {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/specializations?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	specialization := cs.CharacterSpecialization{}
	err = json.NewDecoder(resp.Body).Decode(&specialization)
	return specialization
}

// GenerateCharacterStats method requests characterStats api
func (b *BnetClient) GenerateCharacterStats(name string, realm string) cstats.CharacterStats {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/statistics?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	stats := cstats.CharacterStats{}
	err = json.NewDecoder(resp.Body).Decode(&stats)
	return stats
}

// GenerateCharacterEquipment method requests characterEquipment api
func (b *BnetClient) GenerateCharacterEquipment(name string, realm string) ceq.CharacterEquipment {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	equipment := ceq.CharacterEquipment{}
	err = json.NewDecoder(resp.Body).Decode(&equipment)
	return equipment
}

// GeneratePvP method requests PvP Api giving back arena values
func (b *BnetClient) GeneratePvP(name string, realm string) [2]pvp.CharacterPvp {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/3v3?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}

	returnPvp := [2]pvp.CharacterPvp{}
	arena3v3 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena3v3)
	returnPvp[1] = arena3v3

	requestOut = fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/2v2?namespace=profile-us&locale=en_US", realm, name)
	resp, err = b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	arena2v2 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena2v2)
	returnPvp[0] = arena2v2
	return returnPvp
}

// GeneratePvpSummary method requests PvPSummary Api
func (b *BnetClient) GeneratePvpSummary(name string, realm string) pvps.CharacterPvpSummary {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-summary?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	pvpSummary := pvps.CharacterPvpSummary{}
	err = json.NewDecoder(resp.Body).Decode(&pvpSummary)
	return pvpSummary
}
