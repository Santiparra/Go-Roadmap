package main

func countReports(numSentCh chan int) int {	
	total := 0
	for {
		batch, ok := <- numSentCh
		if !ok {
			break
		}
		total += batch
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

// EXAMPLE 2

func concurrentFib(n int) []int {
	newCh := make(chan int)
	sequence := []int{}
	go fibonacci(n, newCh)
	for num := range newCh {
		sequence = append(sequence, num)
	}
	return sequence
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
