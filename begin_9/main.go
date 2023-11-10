package main

import "fmt"
import "math"

func main() {
	var intA float64
	fmt.Print("Число A:")
	fmt.Scan(&intA)
	
	var intB float64
	fmt.Print("Число B:")
	fmt.Scan(&intB)
	
	average:= math.Sqrt (intA*intB)

	fmt.Println("Среднее геометрическое равняется:", average)


}
