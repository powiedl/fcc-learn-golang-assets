package ch5

import "fmt"

// #region CH5.1
func sendMessage(msg message) (string, int) {
	// ?
	m := msg.getMessage()
	c := len(m)*3
	return m,c
}

type message interface {
	getMessage() string
}
// #endregion

// #region CH5.2
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}
// #endregion

// #region CH5.10
/*
func getExpenseReport(e expense) (string, float64) {
	m, mok := e.(email)
	if mok {
		return m.toAddress, m.cost()
	}
	s, sok := e.(sms)
	if sok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0.0
}
*/

// #CH5.11 - Verbesserung mit type switch
func getExpenseReport(e expense) (string, float64) {
/* my solution
	switch e.(type) {
	case email:
		m, _ := e.(email)
		return m.toAddress, m.cost()
	case sms:
		s, _ := e.(sms)
		return s.toPhoneNumber, s.cost()
	default:
		return "", 0.0
	}
*/
	switch v := e.(type) {
	  case email:
		  return v.toAddress,v.cost()
		case sms:
			return v.toPhoneNumber,v.cost()
		default:
			return "",0.0
	}
}
// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
// #endregion

// #region CH5.C1
type Formatter interface {
	Format() string
}

type Plaintext struct {
	message string
}

func (p Plaintext) Format() string {
	return p.message
}

type Bold struct {
	message string
}

func (b Bold) Format() string {
	return "**" + b.message + "**"
}

type Code struct {
	message string
}

func (c Code) Format() string {
	return "`" + c.message + "`"
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

/*
func main() {
	p := Plaintext{
		message: "Plaintext message",
	}
	b := Bold{
		message: "Bold message",
	}
	c := Code{
		message: "Code message",
	}
	fmt.Printf("Plaintext: %v\n", p.Format())
	fmt.Printf("Bold: %v\n", b.Format())
	fmt.Printf("Code: %v\n", c.Format())
}
*/
// #endregion CH5.C1

// #region CH5.C2
type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

// ?
func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	// ?
	switch n.(type) {
	case directMessage:
		d := n.(directMessage)
		return d.senderUsername, d.importance()
	case groupMessage:
		g := n.(groupMessage)
		return g.groupName, g.importance()
	case systemAlert:
		a := n.(systemAlert)
		return a.alertCode, a.importance()
	default:
		return "", 0
	}
}


/*
func main() {

	d := directMessage{
		senderUsername: "sender",
		messageContent: "message",
		priorityLevel:  7,
		isUrgent:       false,
	}
	a := systemAlert{
		alertCode:      "ERROR",
		messageContent: "Something went wrong",
	}
	strD, impD := processNotification(d)
	fmt.Printf("%v, %v\n", strD, impD)
	strA, impA := processNotification(a)
	fmt.Printf("%v, %v\n", strA, impA)

}
*/
// #endregion

func main() {
	fmt.Println("Chapter 5")

}
