package ch8

import (
	"errors"
	"fmt"
)

type cost struct {
	day   int
	value float64
}


// #region CH8.1
func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgs := [3]string{primary, secondary, tertiary}
	// costs0 := len(primary)
	// costs1 := costs0 + len(secondary)
	// costs2 := costs1 + len(tertiary)
	// return msgs, [3]int{costs0, costs1, costs2}
	// alternativ:
	var costs [3]int
	summedCost := 0
	for i := 0; i < 3; i++ {
		costs[i] = summedCost + len(msgs[i])
		summedCost = costs[i]
	}
	return msgs, costs
}
// #endregion

// #region CH8.2
func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		res := messages[:]
		fmt.Printf("res=%v\n", res)
		return messages[:], nil
	}
	if plan == planFree {
		return messages[:2], nil
	}
	return nil, errors.New("unsupported plan")
}
// #endregion

// #region CH8.6
func getMessageCosts(messages []string) []float64 {
	cost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost[i] = float64(len(messages[i])) * 0.01
	}
	return cost
}
// #endregion

// #region CH8.10
func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
// #endregion

// #region CH8.11

func getDayCosts(costs []cost, day int) []float64 {
	for i := 0; i < len(costs); i++ {
		if day == costs[i].day {
			return []float64{costs[i].value}
		}
	}
	return []float64{}
}
// #endregion

// #region CH8.13

func createMatrix(rows, cols int) [][]int {
	result := [][]int{}
	for r := 0; r < rows; r++ {
		row := make([]int, cols)
		for c := 0; c < cols; c++ {
			row[c] = r * c
		}
		result = append(result, row)
	}
	return result
}
// #endregion

func main() {

	//CH8.14: j and g point to the same underlying array, so g's append overwrote j
	//CH8.15: The array's cap() is exceeded so a new underlying array is allocated
	//CH8.16: always assign the result of the append() function back to the same slice
	
	//CH8.13
	five := createMatrix(5, 5)
	fmt.Printf("%v\n", five)

	//CH8.12
	fmt.Printf("msg: 'this is my message', badWords: 'bad words', %v\n", indexOfFirstBadWord([]string{"this", "is", "my", "message"},[]string{"bad", "words"}))
	fmt.Printf("msg: 'this is my words message', badWords: 'bad words', %v\n",indexOfFirstBadWord([]string{"this","is","my","words","message"},[]string{"bad","words"}));
	//CH8.11
	costs := []cost{
		{day: 1, value: 2.0},
		{day: 2, value: 4.0},
		{day: 3, value: 6.0},
	}
	fmt.Printf("%v\n", getDayCosts(costs, 2))
	fmt.Printf("%v\n", getDayCosts(costs, 4))


	slice3 := make([]float64, 3, 6)
	slice3[0] = 42.0
	fmt.Printf("%v, len: %v, cap: %v \n",slice3,len(slice3),cap(slice3))

	// CH8.10
	s := []int{7, 34, 2, 745, 34, 134, 167}
	fmt.Printf("%v\n", sum(s...))

  // CH8.6
	fmt.Printf("costs: %v\n", getMessageCosts([]string{"hello", "world", "!", "from powidl", "and friends"}))

	// CH8.1
	// msgs, cost := getMessageWithRetries("hello", "world", "powidl")
	fmt.Printf("getMessageWithRetries: %v - messages , %v - costs\n", msgs, cost)

	// CH8.2
	str, err := getMessageWithRetriesForPlan(planFree, [3]string{"hello", "world", "powidl"})
	if err != nil {
		fmt.Printf("%v\n", str)
	} else {
		fmt.Printf("%v\n", err)
	}
	str, err = getMessageWithRetriesForPlan("quaxi", [3]string{"hello", "world", "powidl"})
	if err != nil {
		fmt.Printf("%v\n", str)
	} else {
		fmt.Printf("%v\n", err)
	}
	str, err = getMessageWithRetriesForPlan(planPro, [3]string{"hello", "world", "powidl"})
	if err != nil {
		fmt.Printf("%v\n", str)
	} else {
		fmt.Printf("%v\n", err)
	}
}