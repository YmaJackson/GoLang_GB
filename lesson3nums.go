//lesson3.2
package main

import "fmt"

func main() {
	var number, a, j int
	var Check bool

	fmt.Print("Введите число для вывода всех простых чисел до этого числа: ")
	fmt.Scanln(&number)
	for a = 2; a < number; a++ {

		Check = false
		for j = 2; j <= a/2; j++ {

			if a%j == 0 {
				Check = true
			}
		}
		if Check == false {
			fmt.Println(a)
		}
	}

}
