package ch9

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// #region CH9.1
type user struct {
	name        string
	phoneNumber int
}
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	result := make(map[string]user)
	for i, u := range names {
		us := user{name: u, phoneNumber: phoneNumbers[i]}
		//fmt.Printf("%v:%v => %v\n", u, phoneNumbers[i], us)
		result[u] = us
	}
	return result, nil
}
// #endregion

// #region CH9.2
func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !u.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil
}
// #endregion
// #region CH9.5
func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, u := range messagedUsers {
		if _, ok := validUsers[u]; ok {
			validUsers[u]++
		}
	}
}
// #endregion

// #region CH9.10
func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _, n := range names {
		runes := []rune(n)
		first := runes[0]
		fmt.Printf("first:%v, name:%v\n", first, n)
		if _, ok := result[first]; !ok {
			// es gab noch keinen Namen mit diesem Anfangsbuchstaben, also die Map für diesen Anfangsbuchstaben anlegen
			result[first] = make(map[string]int)
		}
		// Meine Version
		// if _, ok := result[first][n]; !ok {
		// 	// es gab noch keinen Eintrag für diesen Namen, also die Map im richtigen Teilbereich mit 1 anlegen
		// 	result[first][n] = 1
		// } else {
		// 	// es gab schon einen Eintrag für diesen Namen, also den Wert um 1 erhöhen
		// 	result[first][n]++
		// }
		result[first][n]++ // weil wenn n noch nicht existiert, kommt 0 zurück, dass wird um eins erhöt und dann gespeichert
	}
	return result
}
// #endregion

// #region CH9.C1
func countDistinctWords(messages []string) int {
	count := 0
	distinct := map[string]int{}
	for _, m := range messages {
		u := strings.ToUpper(m)
		words := strings.Fields(u)
		for _, w := range words {
			if _, ok := distinct[w]; !ok {
				distinct[w] = 1
				count++
			} else {
				distinct[w]++
			}
		}
	}
	fmt.Printf("%v (distinct words and occurences)\n", distinct)
	return count
}
// #endregion

// #region CH9.C2
func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// ?
	directFriends := friendships[username]
	suggestions := []string{}
	fmt.Printf("%v\n", directFriends)
	for _, dF := range directFriends {
		for _, pF := range friendships[dF] {
			if pF == username {
				// can't be a fried of myself
				continue
			} 
			if slices.Contains(suggestions, pF) {
				// already a possible friend
				continue
			} 
			if !slices.Contains(directFriends, pF) {
				suggestions = append(suggestions, pF)
			}
		}
	}
	return suggestions
}
// #endregion
func main() {
	// CH9.C2
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}

	suggestedFriends := findSuggestedFriends("Alice", friendships)
	fmt.Printf("suggestions: %v\n", suggestedFriends)

	//CH9.C1
	m9C1 := []string{"Hello world", "hello there", "General Kenobi"}
	fmt.Printf("%v distinct words\n", countDistinctWords(m9C1))

  //CH9.10
	n910 := []string{"bob", "billy", "bob", "joe"}
	fmt.Printf("%v\n", getNameCounts(n910))
	
	//CH9.6 - Maps can have at most 1 value(s) associated with the same key
	//CH9.7 - Attempting to get a value from a map where the key doesn't exist... - returns the closest value
	//CH9.8 - changes in a function AFFECT the original values
	//CH9.9 - a boolean that indicates whether the key exists
	//CH9.1
	names := []string{"adam", "bob", "clara"}
	phones := []int{123, 234, 564}
	users, err := getUserMap(names, phones)
	if err != nil {
		fmt.Printf("ERROR:%v\n", err)
	} else {
		fmt.Printf("%v\n", users)
	}

}