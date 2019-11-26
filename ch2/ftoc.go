// Ftoc prints two Fahrenheit-to-Celsius conversions
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g0F = %g0C\n", freezingF, fToC(freezingF)) // "320F = 00C"
	fmt.Printf("%g0F = %g0C\n", boilingF, fToC(boilingF))   // "2120F = 1000C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
