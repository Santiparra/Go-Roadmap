package main

func getNameCounts(names []string) map[rune]map[string]int {
	nested := make(map[rune]map[string]int)
	for _,name := range names {
		rune := []rune(name)[0]
		if _, ok := nested[rune]; !ok {
			nested[rune] = make(map[string]int)
		} 
		nested[rune][name]++
		
	}
	return nested
}