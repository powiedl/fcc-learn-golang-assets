// #region CH14.1
func getLast[T any](s []T) T {
	if len(s) < 1 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}
// #endregion

// CH14.2: Generics help most with a binary tree
// CH14.3: Go only adds new features if they are extremly important (the main goal of Go is to keep it as simple as possible)
// CH14.4: Generics will be used in libraries

// #region CH14.5
func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	cost := newItem.GetCost()
	if cost > balance {
		var r []T
		return r, 0, errors.New("insufficient funds")
	}
	return append(oldItems, newItem), balance - cost, nil
}
// #endregion

// #CH14.6: You know exactly which types satisfy your interface

// #region CH14.7
type biller[C customer] interface {
	Charge(customer C) bill
	Name() string
}
// #endregion

// #CH14.8: The name of a type parameter can be anything, but T is a common convention
