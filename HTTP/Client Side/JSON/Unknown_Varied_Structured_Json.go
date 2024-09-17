package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		return nil, err
	}	
	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for key, value := range resource {
			strigified := fmt.Sprintf("Key: %s - Value: %v", key, value)
			formattedStrings = append(formattedStrings, strigified)
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}