package characterprofile

type CharacterProfile struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"race"`
	CharacterClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_class"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_spec"`
	Realm struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realm"`
	Guild struct {
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
		Faction struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"guild"`
	Level             int `json:"level"`
	Experience        int `json:"experience"`
	AchievementPoints int `json:"achievement_points"`
	Achievements      struct {
		Href string `json:"href"`
	} `json:"achievements"`
	Titles struct {
		Href string `json:"href"`
	} `json:"titles"`
	PvpSummary struct {
		Href string `json:"href"`
	} `json:"pvp_summary"`
	Encounters struct {
		Href string `json:"href"`
	} `json:"encounters"`
	Media struct {
		Href string `json:"href"`
	} `json:"media"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	AverageItemLevel   int   `json:"average_item_level"`
	EquippedItemLevel  int   `json:"equipped_item_level"`
	Specializations    struct {
		Href string `json:"href"`
	} `json:"specializations"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	MythicKeystoneProfile struct {
		Href string `json:"href"`
	} `json:"mythic_keystone_profile"`
	Equipment struct {
		Href string `json:"href"`
	} `json:"equipment"`
	Appearance struct {
		Href string `json:"href"`
	} `json:"appearance"`
	Collections struct {
		Href string `json:"href"`
	} `json:"collections"`
	ActiveTitle struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name          string `json:"name"`
		ID            int    `json:"id"`
		DisplayString string `json:"display_string"`
	} `json:"active_title"`
	Reputations struct {
		Href string `json:"href"`
	} `json:"reputations"`
	Quests struct {
		Href string `json:"href"`
	} `json:"quests"`
	AchievementsStatistics struct {
		Href string `json:"href"`
	} `json:"achievements_statistics"`
	Professions struct {
		Href string `json:"href"`
	} `json:"professions"`
}
