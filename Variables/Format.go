package main

import "fmt"

func Format() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s", fname, lname, age, messageRate, isSubscribed, message)

	fmt.Println(userLog)
}

/*
	%.1f rounds a float to the tenths place, %.2f rounds to the hundredths place, etc.
	%t formats a boolean value.
	%v can be used to format any value in its default representation.
	%s can be used to format a string.
	%d can be used to format an integer.
*/
