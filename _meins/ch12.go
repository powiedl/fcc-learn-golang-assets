package main

import (
	"fmt"
	"time"
)

// #region CH12.1
func sendEmail(message string) {
	go func() {																												// added go
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}
// #endregion

// #CH12.2
func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)
	go sendIsOld(isOldChan, emails) 																	// added go

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// #region CH12.3
func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ { 																		// added this line
		<-dbChan
	}																																	// added this line
}
// #endregion

// #region CH12.4
func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, e := range emails {
		ch <- e
	}
	return ch
}
// #endregion

// #region CH12.5
func countReports(numSentCh chan int) int {
	count := 0
	for {
		current, ok := <-numSentCh
		if !ok {
			break
		}
		count += current
	}
	return count
}
// #endregion

// #region CH12.6
func concurrentFib(n int) []int64 {
	ch := make(chan int64)
	go fibonacci(n, ch)
	result := []int64{}
	for f := range ch {
		result = append(result, f)
	}
	return result
}
// #endregion

// #region CH12.7
func logMessages(chEmails, chSms chan string) {
	select {
	case sms, ok := <-chSms:
		if ok != true {
			return
		}
		logSms(sms)
	case email, ok := <-chEmails:
		if ok != true {
			return
		}
		logEmail(email)
	}
}
// #endregion

// #region CH12.8
// meine Version - funktioniert nicht, nach dem ersten "Nothing to do, waiting..." kommt ien fatal error: all goroutines are asleep - deadlock!
func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
//#endregion

// #region CH12.C1

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, numPings)
	/* go */ func() {																		// removed go - weil sonst die test Funktion beendet wird, bevor pinger und ponger "durch" sind
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
	}()
}
// #endregion

// #region CH12.C2
// mithilfe von Codium - ich bin mir nicht sicher, ob das so im Sinne des Erfinders ist ..., weil es nur funktioniert, weil processMessages weiß, wie viele Nachrichten verarbeitet werden sollen; aber es funktioniert
func processMessages(messages []string) []string {
	result := make(chan string)
	for _, m := range messages {
		go func(m string) {
			processed := process(m)
			result <- processed
		}(m)
	}
	results := []string{}
	for i := 0; i < len(messages); i++ {
		results = append(results, <-result)
	}
	return results
}
// #endregion

// CH12.9: The receiver will block forever
// CH12:10: Panic (Send to a closed channel)


