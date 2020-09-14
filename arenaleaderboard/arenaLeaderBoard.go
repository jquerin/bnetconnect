package arenaleaderboard

type ArenaLeaderBoard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	Name    string `json:"name"`
	Bracket struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"bracket"`
	Entries []struct {
		Character struct {
			Name  string `json:"name"`
			ID    int    `json:"id"`
			Realm struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"realm"`
		} `json:"character"`
		Faction struct {
			Type string `json:"type"`
		} `json:"faction"`
		Rank                  int `json:"rank"`
		Rating                int `json:"rating"`
		SeasonMatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"season_match_statistics"`
		Tier struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"tier"`
	} `json:"entries"`
}
