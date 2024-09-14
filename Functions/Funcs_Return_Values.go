package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	discountRate := calculateDiscountRate(numMessages)
	baseCost := calculateBaseBill(costPerMessage, numMessages)
	if discountRate < 0.1 {
		return baseCost
	}
	return baseCost - (baseCost * discountRate)
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	}
	if messagesSent > 1000 {
		return 0.1
	}
	return 0.0
	
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
