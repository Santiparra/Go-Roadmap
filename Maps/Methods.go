//Insert an element

m[key] = elem

//Get an element

elem = m[key]

//Delete an element

delete(m, key)

//Check if a key exists

elem, ok := m[key]


package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !user.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}