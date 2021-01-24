/*
@author Глобин Андрей
*/
package main

import (
	"fmt"
	"lesson-1/amorphous"
	"lesson-1/other"
	"lesson-1/random"
	"math/rand"
)

func main() {
	// 14. Автоморфные числа
	fmt.Println(amorphous.GenerateNumbers(1000000))

	// 13. случайное число без использования стандартной функции
	fmt.Println(random.GenerateNumber(100))
	// 13. случайное число c использования стандартной функции
	fmt.Println(rand.Intn(100))

	// 12. Написать функцию нахождения максимального из трех чисел.
	fmt.Println(other.Max(100, 105, 101))
}
