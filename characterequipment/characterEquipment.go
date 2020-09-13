package characterequipment

type CharacterEquipment struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
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
	EquippedItems []struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"item"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Quantity  int   `json:"quantity"`
		Context   int   `json:"context,omitempty"`
		BonusList []int `json:"bonus_list,omitempty"`
		Quality   struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Name  string `json:"name"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"media"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_class"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_subclass"`
		InventoryType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Binding struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		Armor struct {
			Value   int `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					R int `json:"r"`
					G int `json:"g"`
					B int `json:"b"`
					A int `json:"a"`
				} `json:"color"`
			} `json:"display"`
		} `json:"armor,omitempty"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value   int `json:"value"`
			Display struct {
				DisplayString string `json:"display_string"`
				Color         struct {
					R int `json:"r"`
					G int `json:"g"`
					B int `json:"b"`
					A int `json:"a"`
				} `json:"color"`
			} `json:"display"`
			IsEquipBonus bool `json:"is_equip_bonus,omitempty"`
		} `json:"stats,omitempty"`
		Spells []struct {
			Spell struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"spell"`
			Description string `json:"description"`
		} `json:"spells,omitempty"`
		Upgrades struct {
			Value         int    `json:"value"`
			MaxValue      int    `json:"max_value"`
			DisplayString string `json:"display_string"`
		} `json:"upgrades,omitempty"`
		Requirements struct {
			Level struct {
				DisplayString string `json:"display_string"`
			} `json:"level"`
		} `json:"omitempty"`
		Level struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"level"`
		TimewalkerLevel int `json:"timewalker_level,omitempty"`
		SellPrice       struct {
			Value          int `json:"value"`
			DisplayStrings struct {
				Header string `json:"header"`
				Gold   string `json:"gold"`
				Silver string `json:"silver"`
				Copper string `json:"copper"`
			} `json:"display_strings"`
		} `json:"sell_price,omitempty"`
		IsSubclassHidden bool `json:"is_subclass_hidden,omitempty"`
		Durability       struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"durability,omitempty"`
		ModifiedAppearanceID int    `json:"modified_appearance_id,omitempty"`
		UniqueEquipped       string `json:"unique_equipped,omitempty"`
		Weapon               struct {
			Damage struct {
				MinValue      int    `json:"min_value"`
				MaxValue      int    `json:"max_value"`
				DisplayString string `json:"display_string"`
				DamageClass   struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"damage_class"`
			} `json:"damage"`
			AttackSpeed struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"attack_speed"`
			Dps struct {
				Value         float64 `json:"value"`
				DisplayString string  `json:"display_string"`
			} `json:"dps"`
		} `json:"weapon,omitempty"`
	} `json:"equipped_items"`
}
