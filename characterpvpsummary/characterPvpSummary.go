package characterpvpsummary

type CharacterPvpSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Brackets []struct {
		Href string `json:"href"`
	} `json:"brackets"`
	HonorLevel       int `json:"honor_level"`
	PvpMapStatistics []struct {
		WorldMap struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"world_map"`
		MatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"match_statistics"`
	} `json:"pvp_map_statistics"`
	HonorableKills int `json:"honorable_kills"`
	Character      struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
}
