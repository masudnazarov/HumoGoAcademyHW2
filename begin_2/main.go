package main

import "fmt"

func main() {

	var a int
	fmt.Println("Чему равна сторона квадрата?")
	fmt.Scan(&a)
	square := a * a
	fmt.Println("Площадь квадрата равна:", square)

}
