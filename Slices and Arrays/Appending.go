package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	resume := []float64{}
	for d := 0; d < len(costs); d++ {		
		for costs[d].day >= len(resume) {
			resume = append(resume, 0.0)
		}
		resume[costs[d].day] += costs[d].value
	}
	return resume
}