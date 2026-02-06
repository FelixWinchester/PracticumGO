package main

import "fmt"

func main() {
    fmt.Println(SumSlice([]int{1, 2, 3, 4, 5}))
    fmt.Println(FindMax([]int{2, 5, 78, 13, 66}))
    fmt.Println(Reverse([]int{1, 2, 3, 4, 5}))
    fmt.Println(CountEven([]int{1, 2, -3, -4, -5, 6, 7, 8, -9, -10}))
    fmt.Println(CreateSequence(3, 7))
    fmt.Println(FindMin([]int{11, 12, 13, 4, 5, 6, 7, 8, 9, 10}))
    fmt.Println(Average([]int{1, 2, 3, 4, 5}))
    fmt.Println(Contains(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
    fmt.Println(MergeSlices([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}))
    fmt.Println(CountGreaterThan([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 10))
}

func SumSlice(slice []int) int {
    result := 0
    for i := 0; i < len(slice); i++ {
        result += slice[i]
    }
    return result
}

func FindMax(slice []int) int {
    if len(slice) == 0 {
        return 0
    }
    max := slice[0]
    for i := 1; i < len(slice); i++ {
        if slice[i] > max {
            max = slice[i]
        }
    }
    return max
}

func Reverse(slice []int) []int {
    result := make([]int, len(slice))
    for i := 0; i < len(slice); i++ {
        result[i] = slice[len(slice)-1-i]
    }
    return result
}

func CountEven(slice []int) int {
    counter := 0
    for i := 0; i < len(slice); i++ {
        if slice[i]%2 == 0 {
            counter++
        }
    }
    return counter
}

func CreateSequence(start int, end int) []int {
    result := make([]int, 0, end-start+1)
    for i := start; i <= end; i++ {
        result = append(result, i)
    }
    return result
}

func FindMin(slice []int) int {
    if len(slice) == 0 {
        return 0
    }
    min := slice[0]
    for i := 1; i < len(slice); i++ {
        if slice[i] < min {
            min = slice[i]
        }
    }
    return min
}

func Average(slice []int) float64 {
    if len(slice) == 0 {
        return 0
    }
    sum := 0
    for i := 0; i < len(slice); i++ {
        sum += slice[i]
    }
    return float64(sum) / float64(len(slice))
}

func Contains(value int, slice []int) bool {
    for i := 0; i < len(slice); i++ {
        if slice[i] == value {
            return true
        }
    }
    return false
}

func MergeSlices(slice1 []int, slice2 []int) []int {
    result := make([]int, len(slice1)+len(slice2))
    copy(result, slice1)
    copy(result[len(slice1):], slice2)
    return result
}

func CountGreaterThan(slice []int, n int) int {
    counter := 0
    for i := 0; i < len(slice); i++ {
        if slice[i] > n {
            counter++
        }
    }
    return counter
}
