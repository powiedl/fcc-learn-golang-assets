package main

import (
	"fmt"
	"testing"
	"time"
)

func processMessages(messages []string) []string {
	result := []string{}
	ch := make(chan string, 3)
	//	go func() {
	//	}()
	for _, m := range messages {
		go func() {
			ch <- process(m)
		}()
	}
	for {
		m, ok := <-ch
		if !ok {
			return result
		}
		result = append(result, m)
	}

}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

func TestProcessMessages(t *testing.T) {
	type testCase struct {
		messages []string
		expect   []string
	}

	runCases := []testCase{
		{
			messages: []string{"Sunlit", "Man"},
			expect:   []string{"Man-processed", "Sunlit-processed"},
		},
		{
			messages: []string{"Nomad do you copy?"},
			expect:   []string{"Nomad do you copy?-processed"},
		},
		{
			messages: []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"},
			expect:   []string{"Taldain-processed", "Roshar-processed", "Sel-processed", "Nalthis-processed", "Scadriel-processed"},
		},
	}

	submitCases := append(runCases, []testCase{
		{
			messages: []string{},
			expect:   []string{},
		},
		{
			messages: []string{"Scadriel"},
			expect:   []string{"Scadriel-processed"},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		fail := false
		result := processMessages(test.messages)

		if len(result) != len(test.expect) {
			fail = true
		}

		counts := make(map[string]int)
		for _, res := range result {
			counts[res]++
		}
		for _, exp := range test.expect {
			counts[exp]--
			if counts[exp] < 0 {
				fail = true
			}
		}

		if fail {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
  inputs:   %v
  expected: %v
  actual:   %v
  `, test.messages, test.expect, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  inputs:   %v
  expected: %v
  actual:   %v
`, test.messages, test.expect, result)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true