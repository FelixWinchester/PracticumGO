package main

import "fmt"

func main() {
	fmt.Println(MinMax([]int{10, 3, 8, 1, 9}))
	fmt.Println(FindIndex(15, []int{5, 10, 15, 20, 25}))
	fmt.Println(CountPositive([]int{-5, 10, 0, 3, -2, 7}))
	fmt.Println(MultiplyBy(3, []int{1, 2, 3, 4, 5}))
	fmt.Println(SumTwoSlices([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ReplaceElement(2, 99, []int{1, 2, 3, 2, 5, 2, 6}))
	fmt.Println(FirstN(3, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(InterLeave([]int{1, 2, 3}, []int{10, 20, 30}))
}

func MinMax(slice1 []int) int {
	if len(slice1) == 0 {
		return 0
	}
	max := slice1[0]
	min := slice1[0]
	for i := 0; i < len(slice1); i++ {
		if slice1[i] > max {
			max = slice1[i]
		}
		if slice1[i] < min {
			min = slice1[i]
		}
	}
	return max - min
}

func FindIndex(n int, slice2 []int) int {
	for i, v := range slice2 {
		if v == n {
			return i
		}
	}
	return -1
}

func CopySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		copySlice[i] = slice[i]
	}
	return copySlice
}

func IsPalindrome(slice []int) bool {
	for i := 0; i < len(slice)/2; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}

func CountPositive(slice5 []int) int {
	counter := 0
	for i := 0; i < len(slice5); i++ {
		if slice5[i] > 0 {
			counter++
		}
	}
	return counter
}

func MultiplyBy(n int, slice6 []int) []int {
	answer := []int{}
	for i := 0; i < len(slice6); i++ {
		answer = append(answer, slice6[i]*n)
	}
	return answer
}

func SumTwoSlices(slice7_1 []int, slice7_2 []int) []int {
	if len(slice7_1) != len(slice7_2) {
		return []int{}
	}
	sumSlices := []int{}
	for i := 0; i < len(slice7_1); i++ {
		sumSlices = append(sumSlices, slice7_1[i]+slice7_2[i])
	}
	return sumSlices
}

func ReplaceElement(n int, m int, slice8 []int) []int {
	for i := 0; i < len(slice8); i++ {
		if slice8[i] == n {
			slice8[i] = m
		}
	}
	return slice8
}

func FirstN(n int, slice9 []int) []int {
	if n > len(slice9) {
		n = len(slice9)
	}
	answer := slice9[:n]
	return answer
}

func InterLeave(slice10_1 []int, slice10_2 []int) []int {
	answer := []int{}
	for i := 0; i < len(slice10_1); i++ {
		answer = append(answer, slice10_1[i])
		answer = append(answer, slice10_2[i])
	}
	return answer
}
