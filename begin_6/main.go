package main

import "fmt"

func main() {
	var edgeA int
	fmt.Print("Чему равна длина ребра A?")
	fmt.Scan(&edgeA)

	var edgeB int
	fmt.Print("Чему равна длина ребра B?")
	fmt.Scan(&edgeB)

	var edgeC int
	fmt.Print("Чему равна длина ребра C?")
	fmt.Scan(&edgeC)

	volume := edgeA * edgeB * edgeC
	square := 2 * (edgeA*edgeB + edgeB*edgeC + edgeA*edgeC)
	fmt.Println("Объём прямоугольника состовляет:", volume, "Площадь поверхности состовляет:", square)

}
