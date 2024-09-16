package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var sliOfsli [][]byte
	for _, item := range items {		
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err	
		}
		sliOfsli = append(sliOfsli, data)
	}		
	return sliOfsli, nil
}
