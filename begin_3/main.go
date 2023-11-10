package main

import "fmt"

func main() {

	var a int
	fmt.Println("Каков сторона a?")
	fmt.Scan(&a)

	var b int
	fmt.Println("Какова сторона b")
	fmt.Scan(&b)

	s := a * b
	perimeter := 2 * (a + b)
	fmt.Println("Площадь квадрата:", s, "Периметр равен:", perimeter)

}
