package main

import "math/big"

func main() {
	var a *big.Int
	a.Exp(big.NewInt(10), big.NewInt(99), nil)

	println(a)
	println(a.ProbablyPrime(20))
}
