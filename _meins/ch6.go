package ch6

import (
	"errors"
	"fmt"
)

// #region CH6.1
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	// ?
	costCust, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	costSpouse, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err2
	}
	return costCust + costSpouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
// #endregion

// #region CH6.2
func getSMSErrorString(cost float64, recipient string) string {
	// ?
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",cost,recipient)
}
// #endregion

// #region CH6.3
type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero",d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
// #endregion

// #region CH6.6
func divide66 (dividend,divisor float64) (float64,error) {
	if divisor == 0 {
		return 0,errors.New("no dividing by 0")
	}
	return dividend/divisor,nil
}
// #endregion

// #region CH6.C1
func validateStatus(status string) error {
	if status == "" { return errors.New("status cannot be empty")}
	if len(status)>140 { return errors.New("status exceeds 140 characters")}
	return nil
}
// #endregion
func main() {
	// CH6.1
	c1, e1 := sendSMSToCouple("bla", "blub")
	c2, e2 := sendSMSToCouple("ba", "blubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblub")
	c3, e3 := sendSMSToCouple("blubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblubblub", "bla")
	if e1 == nil {
		fmt.Printf("Sending SMS, total cost: %v\n", c1)
	} else {
		fmt.Printf("Something went wrong %v\n", e1)
	}
	if e2 == nil {
		fmt.Printf("Sending SMS, total cost: %v\n", c2)
	} else {
		fmt.Printf("Something went wrong:%v\n", e2)
	}
	if e3 == nil {
		fmt.Printf("Sending SMS,total cost: %v\n", c3)
	} else {
		fmt.Printf("Something went wrong %v\n", e3)
	}

	// CH6.2
	fmt.Println(getSMSErrorString(3.42342,"powidl"))

	// CH6.3
	r,err := divide(7.0,3.0)
	if err == nil {fmt.Printf("Result %.3f",r) } else {fmt.Println(err)}
	r,err = divide(7.0,0)
	if err == nil {fmt.Printf("Result %.3f",r) } else {fmt.Println(err)}
	r,err := divide66(7.0,3.0)
	if err == nil {fmt.Printf("Result %.3f",r) } else {fmt.Println(err)}
	r,err = divide66(7.0,0)
	if err == nil {fmt.Printf("Result %.3f",r) } else {fmt.Println(err)}

	// CH6.C1
	err = validateStatus("")
	fmt.Printf("validateStatus returns %v\n", err)
	err = validateStatus("fashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djklfashjfh asdjkfd safhd fahf jskaf hajkfl ah fjkldsahf asdjkfh a fjk hdkasjlf dhjfk fhjk dsfhjkasd fhak flhsf jkfhasd djkl")
	fmt.Printf("validateStatus returns %v\n", err)
	err = validateStatus("happy")
	fmt.Printf("validateStatus returns %v\n", err)

}


