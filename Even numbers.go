/*

Темы: функции, слайсы, циклы, defer
Напишите программу, которая:

    Создаёт слайс чисел от 1 до N (N вводится пользователем).

    Напишите функцию filterEven(arr []int) []int, возвращающую слайс только чётных чисел.

    В main используйте defer для вывода оригинального и отфильтрованного слайсов.

    Добавьте простой пользовательский ввод для N.

*/

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите количество элементов (N): ")
	fmt.Scan(&n)

	allNumb := make([]int, n)

	fmt.Println("Введите элементы: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&allNumb[i])
	}

	defer fmt.Println("Ваш исходный слайс: ", allNumb)
	fmt.Println("")
	fmt.Println("Ваш слайс четных чисел: ", filterEven(allNumb))
}

func filterEven(arr []int) []int {
	filtered := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			filtered = append(filtered, arr[i])
		}
	}
	return filtered
}
