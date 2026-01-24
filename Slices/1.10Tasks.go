package main

import (
	"fmt"
)

func main() {
	// 1 task
	task1 := [3]string{"Apple", "banana", "orange"}
	fmt.Println("Task1: ", task1)

	fmt.Println(" ")

	// 2 task
	task2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Task2: ", task2, len(task2), task2[2])

	fmt.Println(" ")

	// 3 Task
	task3 := []int{10, 20, 30}
	task3 = append(task3, 40)
	fmt.Println("Task3: ", task3)

	fmt.Println(" ")

	// 4 Task
	task4 := []int{}
	task4 = append(task4, 10, 15, 20)
	fmt.Println("Task4: ", task4)

	fmt.Println(" ")

	// 5 Task
	task5 := []string{"A", "B", "C", "D", "E"}
	task5 = task5[1:4]
	fmt.Println("Task5: ", task5)

	fmt.Println(" ")

	// 6 Task
	task6 := []int{100, 200, 300}
	slice := task6[0:2]
	slice[0] = 999
	fmt.Println("Task6: ", task6, slice)

	fmt.Println(" ")

	// 7 Task
	task7 := []string{"красный", "зеленый", "синий", "желтый"}
	fmt.Println("Task7: ")
	for i, v := range task7 {
		fmt.Println(i, v)
	}

	fmt.Println(" ")

	// 8 Task
	task8 := []int{7, 8, 9}
	copy1 := make([]int, len(task8))
	copy(copy1, task8)
	fmt.Println("Task8: ", task8, copy1)

	fmt.Println(" ")

	// 9 Task
	task9 := []float64{1.5, 2.5, 3.0}
	summ := []float64{}
	var iteration float64

	for i := 0; i < len(task9); i++ {
		iteration += task9[i]
	}

	summ = append(summ, iteration)
	fmt.Println("Task9: ", summ)

	fmt.Println(" ")

	// 10 Task
	part1 := []int{1, 2}
	part2 := []int{3, 4, 5}
	part3 := []int{6}

	allNumbers := []int{}
	allNumbers = append(allNumbers, part1...)
	allNumbers = append(allNumbers, part2...)
	allNumbers = append(allNumbers, part3...)

	fmt.Println("Task10: ", allNumbers)
	fmt.Println("Длина: ", len(allNumbers))
}
