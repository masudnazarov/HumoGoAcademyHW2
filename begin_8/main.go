package main

import "fmt"

func main() {
	var intA int
	fmt.Print("Число A:")
	fmt.Scan(&intA)

	var intB int
	fmt.Print("Число B:")
	fmt.Scan(&intB)

	average:=(intA+intB)/2
	fmt.Println("Среднее арифметическое равняется:", average)
	
}
