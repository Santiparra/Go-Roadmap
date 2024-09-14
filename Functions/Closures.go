package main

func adder() func(int) int {
	sum := 0
	return func (num int) int {
		sum += num
		return sum
	}
}
