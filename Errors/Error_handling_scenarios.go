package main

import (
	"errors"
)

/*
 log.Fatal to print a message and exit the program.
*/

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

///////////////////////////////////////////////////////

func enrichUser(userID string) (User, error) {
	user, err := getUser(userID)
	if err != nil {
			return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

//////////////////////////////////////////////////////////

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
