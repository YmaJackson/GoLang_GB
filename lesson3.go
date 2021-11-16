package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, %,  pov, sqrt, cos, sin, log): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("на ноль делить нельзя")
		}
		res = a / b

	case "pov":
		res = math.Pow(a, b)
	case "sqrt2":
		res = math.Sqrt(a)
	case "cos":
		res = math.Cos(a)
	case "sin":
		res = math.Sin(a)

	case "log":
		res = math.Log(a)

	case "%":
		res = math.Mod(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
