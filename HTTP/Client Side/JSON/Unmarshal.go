package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var items []Item
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil 
}