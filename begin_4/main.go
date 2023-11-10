package main

import "fmt"

func main() {

	var diameter float64
	fmt.Println("Чему равен диаметр?")
	fmt.Scan(&diameter)
	length := diameter * 3.14
	fmt.Println("Длина окружности состовляет:", length)

}
