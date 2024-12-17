package classical_150

// canCompleteCircuit 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	minBalance, minIndex, balance := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		balance += gas[i] - cost[i]
		if balance < minBalance {
			minBalance = balance
			minIndex = i + 1
		}
	}
	if balance >= 0 {
		return minIndex
	}
	return -1
}
