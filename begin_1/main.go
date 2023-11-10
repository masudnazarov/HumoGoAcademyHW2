package main

import "fmt"

func main() {

	var part int

	fmt.Println("Чему равна сторона квадрата?")
	fmt.Scan(&part)
	perimeter := part * 4
	fmt.Println("Периметр квадрата равен: ", perimeter)

}
