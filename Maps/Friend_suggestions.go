package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	directFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		directFriends[friend] = true
	}

	mutualFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !directFriends[friendOfFriend] {
				mutualFriends[friendOfFriend] = true
			}
		}
	}

	var suggestions []string
	for friend := range mutualFriends {
		suggestions = append(suggestions, friend)
	}

	return suggestions
}