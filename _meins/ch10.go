package ch10

import (
	"errors"
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m *Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, (*m).Recipient, (*m).Text)
}

// #region CH10.2 und CH10.6
type replace struct {
	word        string
	replacement string
}
func removeProfanity(message *string) {
	if message == nil {
		fmt.Println("EMTPY Message, my wisdom prevented a PANIC")
		return
	}
	// ?
	replacements := []replace{
		{word: "fubb", replacement: "****"},
		{word: "shiz", replacement: "****"},
		{word: "witch", replacement: "*****"},
	}
	for _, e := range replacements {
		*message = strings.ReplaceAll(*message, e.word, e.replacement)
	}
}
// #endregion

// #region CH10.3
type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message103 struct {
	Recipient string
	Success   bool
}

func analyse(analytics *Analytics, m Message103) {
	if m.Success {
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesFailed++
	}
	analytics.MessagesTotal++
}
// #endregion

// #region CH10.8
type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}
// #endregion

// #region CH10.C1
type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cust *customer, t transaction) error {
	if (t.transactionType != transactionDeposit) && (t.transactionType != transactionWithdrawal) {
		return errors.New("unknown transaction type")
	}
	if t.transactionType == transactionWithdrawal {
		if cust.balance < t.amount {
			return errors.New("insufficient funds")
		}
		cust.balance -= t.amount
	} else {
		cust.balance += t.amount
	}
	return nil
}
// #endregion
func main() {
	//CH10.4 - *y = 100
	//CH10.5 - x = 100
	//CH10.7 - pointer receivers are more widely used in Go
	//CH10.9 - Pointer usually points to Heap
	//CH10.10 - Pass by value is usually faster (especially for small amounts of data)

	//CH10.C1
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}
	err := updateBalance(&alice, deposit)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	err = updateBalance(&alice, transaction{customerID: 1, amount: 90, transactionType: transactionWithdrawal})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	err = updateBalance(&alice, transaction{customerID: 1, amount: 90, transactionType: transactionWithdrawal})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Printf("%v\n", alice)


	//CH10.8
	em := email{
		message:     "old",
		fromAddress: "from",
		toAddress:   "to",
	}
	em.setMessage("my new message")
	fmt.Println(em.message) // "my new message"	

	//CH10.3
	a := Analytics{}
	analyse(&a, Message103{Recipient: "r1", Success: true})
	analyse(&a, Message103{Recipient: "r2", Success: false})
	fmt.Printf("analytics: %v\n", a)

	//CH10.2
	s := "bla fubbblu bla blushiz bla witchwitch bla"
	removeProfanity(&s)
	fmt.Printf("%v\n", s)

	//CH10.1
	m := Message{
		Recipient: "empfaenger",
		Text:      "Nachrichtentext",
	}
	fmt.Printf("%v\n", getMessageText(&m))
}
