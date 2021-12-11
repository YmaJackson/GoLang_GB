package main

import (
	"fmt"
	"math"
	"os"
)

// Структура значений вводных данных
type Value struct {
	a float64
	b float64
}

func (s *Value) Addition() float64 {
	return s.a + s.b
}

func (s *Value) Subtraction() float64 {
	return s.a - s.b
}

func (s *Value) Multiplication() float64 {
	return s.a * s.b
}

func (s *Value) Division() float64 {
	return s.a / s.b
}

func (s *Value) Squareroot() float64 {
	return math.Sqrt(float64(s.a))
}
func (s *Value) cos() float64 {
	return math.Cos(s.a)
}
func (s *Value) sin() float64 {
	return math.Sin(s.a)
}
func (s *Value) Log() float64 {
	return math.Log(s.a)
}

type CalculatedValue interface {
	Addition() float64
	Subtraction() float64
	Multiplication() float64
	Division() float64
	Squareroot() float64
	cos() float64
	sin() float64
	Log() float64
}

func main() {
	var d, e, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&d)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&e)

	fmt.Print("Введите операцию (+, -, *, /, sqrt, cos, sin, log, quit): ")
	fmt.Scanln(&op)
	val := &Value{a: d, b: e}

	switch op {
	case "+":
		res = CalculatedValue.Addition(val)
	case "-":
		res = CalculatedValue.Subtraction(val)
	case "*":
		res = CalculatedValue.Multiplication(val)
	case "/":
		if e == 0 {
			fmt.Println("Ошибка на ноль делить нельзя")
			return
		}
		res = CalculatedValue.Division(val)
	case "sqrt":
		res = CalculatedValue.Squareroot(val)

	case "cos":
		res = CalculatedValue.cos(val)
	case "sin":
		res = CalculatedValue.sin(val)
	case "log":
		res = CalculatedValue.Log(val)

	case "quit":
		os.Exit(1)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
