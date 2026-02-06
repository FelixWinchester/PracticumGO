package main

import (
	"fmt"
	"slices"
)

func main() {
	// 1
	fmt.Println(SumSlice([]int{1, 2, 3, 4, 5}))
	// 2
	fmt.Println(FindMaxx([]int{2, 5, 78, 13, 66}))
	// 3
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5}))
	// 4
	fmt.Println(CountEven([]int{1, 2, -3, -4, -5, 6, 7, 8, -9, -10}))
	// 5
	fmt.Println(CreateSequence(3, 7))
	// 6
	fmt.Println(FindMin([]int{11, 12, 13, 4, 5, 6, 7, 8, 9, 10}))
	// 7
	fmt.Println(Average([]int{1, 2, 3, 4, 5}))
	//8
	fmt.Println(Countains(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// 9
	fmt.Println(MergeSlices([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}))
	// 10
	fmt.Println(CountGreaterThan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 10))
}

func SumSlice(slice []int) int {
	result := 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	return result
}

func FindMaxx(slice2 []int) int {
	max := 0
	for i := 0; i < len(slice2); i++ {
		if slice2[i] > max {
			max = slice2[i]
		}
	}
	return max
}

func Reverse(slice3 []int) []int {
	for i, j := 0, len(slice3)-1; i < j; i, j = i+1, j-1 {
		slice3[i], slice3[j] = slice3[j], slice3[i]
	}
	return slice3
}

func CountEven(slice4 []int) int {
	counter := 0
	for i := 0; i < len(slice4); i++ {
		if slice4[i]%2 == 0 {
			counter++
		}
	}
	return counter
}

func CreateSequence(n int, m int) []int {
	slice5 := []int{}
	for i := n; i <= m; i++ {
		slice5 = append(slice5, i)
	}
	return slice5
}

func FindMin(slice6 []int) int {
	result := slices.Min(slice6)
	return result
}

func Average(slice7 []int) float64 {
	sum := 0
	for i := 0; i < len(slice7); i++ {
		sum += slice7[i]
	}
	return float64(sum / len(slice7))
}

func Countains(n int, slice8 []int) bool {
	return slices.Contains(slice8, n)
}

func MergeSlices(slice9_1 []int, slice9_2 []int) []int {
	slice9_1 = append(slice9_1, slice9_2...)
	return slice9_1
}

func CountGreaterThan(slice10 []int, n int) int {
	counter := 0
	for i := 0; i < len(slice10); i++ {
		if slice10[i] > n {
			counter++
		}
	}
	return counter
}
