package main

import (
	"fmt"
	"math/big"
	"os"
)

/*
22.	Разработать программу,
которая перемножает,
делит,
складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.

1) BigInt
2) самостоятельная реализация
*/

func main() {

	var val1Str, val2Str, action string
	fmt.Println("Enter first value:")
	_, err := fmt.Fscan(os.Stdin, &val1Str)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter second value:")
	_, err = fmt.Fscan(os.Stdin, &val2Str)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter acton (+, -, *, /)")
	_, err = fmt.Fscan(os.Stdin, &action)
	if err != nil { //|| action != "+" || action != "-" || action != "*" || action != "/" {
		panic(err)
	}

	n := new(big.Int)
	val1Big, ok := n.SetString(val1Str, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	n2 := new(big.Int)
	val2Big, ok := n2.SetString(val2Str, 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}

	switch action {
	case "*":
		res := MultiplyBigInt(val1Big, val2Big)
		fmt.Printf("\nResult of MUL %d on %d: %v", val1Big, val2Big, res.String())
	case "/":
		res := DivideBigInt(val1Big, val2Big)
		fmt.Printf("\nResult of DIV %d on %d: %v", val1Big, val2Big, res.String())
	case "+":
		res := AddBigInt(val1Big, val2Big)
		fmt.Printf("\nResult of ADD %d on %d: %v", val1Big, val2Big, res.String())
	case "-":
		res := SubtractBigInt(val1Big, val2Big)
		fmt.Printf("\nResult of SUB %d on %d: %v", val1Big, val2Big, res.String())
	}

}

// Multiplication
func MultiplyBigInt(val1 *big.Int, val2 *big.Int) (res big.Int) {
	return *res.Mul(val1, val2)
}

// Divide
func DivideBigInt(val1 *big.Int, val2 *big.Int) (res big.Int) {
	return *res.Div(val1, val2)
}

// Add
func AddBigInt(val1 *big.Int, val2 *big.Int) (res big.Int) {
	return *res.Add(val1, val2)
}

// Subtract
func SubtractBigInt(val1 *big.Int, val2 *big.Int) (res big.Int) {
	return *res.Sub(val1, val2)
}
