package main

import (
	"fmt"
)

func lootTreasure(baseURL, chestRarity string) []Item {
	amountOfItems := map[string]int{
		"Common":    1,
		"Rare":      3,
		"Legendary": 5,
	}

	fullURL := fmt.Sprintf("%s?sort=quality&limit=%d", baseURL, amountOfItems[chestRarity])
	return getItems(fullURL)
}