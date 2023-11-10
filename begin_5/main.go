package main

import "fmt"

func main() {
	var edge int
	fmt.Print("Чему равна длина ребра? ")
	fmt.Scan(&edge)
	volume := edge ^ 3
	square := 6* edge ^ 2
	fmt.Println("Объём куба состовляет:", volume, "Площадь поверхности равна:", square)

}
