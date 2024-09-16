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