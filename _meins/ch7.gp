// You can edit this code!
// Click here and start typing.
package ch7

import (
	"fmt"
)

// #region CH7.1
func bulkSend(numMessages int) float64 {
	sum := 0.0
	for i := 0; i < numMessages; i++ {
		cost := 1.0 + float64(i)/100
		sum = sum + cost
	}
	return sum
}
// #endregion

// #region CH7.2
func maxMessages(thres int) int {
	sum := 0
	for i := 0; ; i++ {
		cost := 100 + i
		sum = sum + cost
		if sum > thres {
			return i
		}
	}
}
// #endregion

// #region CH7.3
func fizzbuzz() {
	for i := 1; i <= 20; i++ {
		s := ""
		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%v\n", i)
		} else {
			if i%3 == 0 {
				s = "fizz"
			}
			if i%5 == 0 {
				s = s + "buzz"
			}
			fmt.Printf("%v\n", s)
		}

	}
}

// #region CH7.5
func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		
		p := true
		for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
//  for i := 3; i * i<= n; i++ {     // Alternativer Ansatz - solange i^2 <= n ist ...
			fmt.Printf("  %v as divisor:", i)
			if n%i == 0 {
				p = false
				break
			}
		}
		if p {
			fmt.Println(n)
		}
	}
}
// #endregion

// #region CH7.C1
func countConnections(groupSize int) int {
	conn := 0
	for i := 2; i <= groupSize; i++ {
		conn = conn + i - 1
	}
	return conn
}
// #endregion

func main() {
	// CH7.1
  fmt.Printf("Der Vesand von %v Message(s) kostet %.2f\n", 1, bulkSend(1))
	fmt.Printf("Der Vesand von %v Message(s) kostet %.2f\n", 5, bulkSend(5))
	fmt.Printf("Der Vesand von %v Message(s) kostet %.2f\n", 10, bulkSend(10))

  // CH7.2
  fmt.Printf("Mit %v Cent kann man max %v Messages schicken\n", 9990, maxMessages(9990))

  // CH7.3
  fizzbuzz()

  // CH7.5
  printPrimes(10)

  // CH7.C1
	fmt.Printf("%v Member haben %v connections\n", 2, countConnections(2))
	fmt.Printf("%v Member haben %v connections\n", 3, countConnections(3))
	fmt.Printf("%v Member haben %v connections\n", 5, countConnections(5))

}
