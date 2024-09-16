package main

func getLast[T any](s []T) T {
	var myZero T
	lenght := len(s)
	if lenght < 1 {
		return myZero
	}
	return s[lenght -1]
	
}
