package ch8

import (
	"errors"
	"fmt"
	"strings"
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
		return messages[:], nil
		return allMessages, nil
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

// #region CH8.C1
func filterMessages(messages []Message, filterType string) []Message {
	result := []Message{}
	for _, el := range messages {
		if el.Type() == filterType {
			result = append(result, el)
		}
	}
	return result
}
// #endregion

// #region CH8.C2
func isValidPassword(password string) bool {
	uppercase, digit := false, false
	uPassword := strings.ToUpper(password)
	fmt.Printf("Uppercase Password:%v\n", uPassword)
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	for i, el := range password {
		if strings.ContainsRune("0123456789", el) {
			digit = true
		} else {
			// Muss ins else des Ziffernchecks, weil bei einer Ziffer ist der Uppercase string auch gleich!
			if el == rune(uPassword[i]) {
				uppercase = true
			}
		}
		if uppercase && digit {
			return true
		}
	}
	return false
}
// #endregion

// #region CH8.C3 - nur die Funktionssignatur war zu korrigieren
func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}
// #endregion

// #region CH8.C4 - der Rest ist in der Angabe
func printReports(messages []string) {
	// ?
	for _, m := range messages {
		printCostReport(func(m string) int {
			return len(m)
		}, m)
	}
}
// #endregion

// #region CH8.C5
func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	result := []sms{}
	for _, m := range messages {
		t := tagger(m)
		result = append(result, sms{id: m.id, content: m.content, tags: t})
	}
	return result
}

func tagger(msg sms) []string {
	tags := []string{}
	uMsg := strings.ToUpper(msg.content)
	if strings.Contains(uMsg, "URGENT") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(uMsg, "SALE") {
		tags = append(tags, "Promo")
	}
	return tags
}
// #endregion

// #region CH8.C6
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(first, second string) {
		s := formatter(first, second)
		fmt.Print(s)
	}
}
// #endregion


func main() {
	// CH8.C5
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		{id: "003", content: "Last saLe! So it is urgenT!"},
		{id: "004", content: "fa era fh lkj2354jkr ajkf er afhjg fghj"},
	}
	fmt.Printf("%v\n", tagMessages(messages, tagger))

	// CH8.C2
	fmt.Printf("true=%v\n", isValidPassword("12345adf%D"))
	fmt.Printf("false=%v\n", isValidPassword("1234"))
	fmt.Printf("false=%v\n", isValidPassword("123456aZ7898012"))
	fmt.Printf("false=%v\n", isValidPassword("fer34afda"))
	fmt.Printf("false=%v\n", isValidPassword("afsaFASDFD"))	
	// CH8.C1
	mm1 := MediaMessage{Sender: "ms1", MediaType: "mmt1", Content: "mc1"}
	mm2 := MediaMessage{Sender: "ms2", MediaType: "mmt2", Content: "mc2"}
	tm1 := TextMessage{Sender: "ts", Content: "tc"}
	lm1 := LinkMessage{Sender: "ls", URL: "lu", Content: "lc"}
	messages := []Message{mm1, mm2, tm1, lm1}
	fmt.Printf("%v\n", filterMessages(messages, "link"))


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