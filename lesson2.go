//lesson 2
package main

import (
	"fmt"
	"math"
)

func main() {

	// 1 task - вычисление площади прямоугольника, длины сторон должны должны вводиться с клавиатуры
	var a, b, S float64
	var C, hundreds, dozens, units int
	const pi = math.Pi
	fmt.Println("введите длину")
	fmt.Scanln(&a)

	fmt.Println("введитее ширину")
	fmt.Scanln(&b)

	fmt.Printf("площадь прямоугольника : %f\n", a*b)

	// 2 - вычисление диаметра и длины окружности по заданной площади круга, площадь круга вводить с клавиатуры

	fmt.Println("введите площадь круга")
	fmt.Scanln(&S)

	fmt.Printf("Диаметр равен : %f\n", (math.Sqrt(S*4*pi))/pi)
	fmt.Printf("длина окружности равна : %f\n", math.Sqrt(S*4*pi))

	// 3 - с клавиатуры вводится 3-х значное число, вывести цифры соответствующие количеству сотен,десятков, единиц
	fmt.Printf("введите трехзначное число")
	fmt.Scanln(&C)
	hundreds = C / 100
	units = C % 10
	dozens = (C % 100) / 10
	fmt.Println(hundreds, "- сотен", dozens, "- десятков", units, "- единиц")

}
