package main

import (
	"fmt"

	"time"
)

type cache struct {
	Map map[int]int
}

var (
	m = cache{Map: map[int]int{}}
)

func main() {

	for {

		var n int

		fmt.Print("Введите число: ")

		fmt.Scanln(&n)

		startAt := time.Now()

		fmt.Println("время исполнения")

		fmt.Println(fibMapCache(n, m), "= результат от числа", n)

		fmt.Println("Проверка cashe : ", time.Now().Sub(startAt))
		startAt = time.Now()
		fmt.Println(fibbonachi(n), "= результат от числа", n)
		fmt.Println("Проверка recursion: ", time.Now().Sub(startAt))

	}

}

func fibbonachiMapCache(x int, m cache) int {

	if x < 2 {

		m.Map[x] = x
		return x
	}

	_, ok := m.Map[x-1]

	if !ok {

		m.Map[x-1] = fibbonachiMapCache(x-1, m)
	}

	return m.Map[x-1] + m.Map[x-2]

}

func fibbonachi(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)

}
