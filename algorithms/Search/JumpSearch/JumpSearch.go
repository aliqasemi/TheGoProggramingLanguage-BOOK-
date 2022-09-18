package main

import (
	"fmt"
	"math"
	"time"
)

func binarySearch(arr []int, value int, defaultIndex ...int) int {
	var defaultIndexValue int   //	this variable store defaultIndex from argument
	if len(defaultIndex) == 0 { //	check if defaultIndex parameter does not exist set defaultIndexValue is 0
		defaultIndexValue = 0
	} else {
		defaultIndexValue = defaultIndex[0] //	check if defaultIndex is existed, set defaultIndexValue is first index of parameter
	}
	mid := len(arr) / 2 // divide array two parts

	if arr[mid] == value {
		return mid + defaultIndexValue
		/*	if find value, return index of value (min) plus defaultIndexValue,
			because defaultIndexValue illustrate last min var index from last recursive func	*/
	} else if mid < 1 { // value does not exist in arr
		return -1
	} else if value < arr[mid] {
		return binarySearch(arr[:mid], value, defaultIndexValue) // search on right side
	} else {
		defaultIndexValue += mid
		return binarySearch(arr[mid:], value, defaultIndexValue) //	search on left side, consider defaultIndexValue for preserve index min from last recursive func
	}
}

func jumpSearch(arr []int, find int) int {
	step := int(math.Sqrt(float64(len(arr)))) // define step for jump index array
	index := 0
	for index < len(arr)-1 { // while condition for jump array
		index = (index + step) - ((index+step)/(len(arr)-1))*((index+step)%(len(arr)-1)) //	update index from step value and if value more than length of array, set value for length of array
		if find <= arr[index] {
			return binarySearch(arr[(index-step):index+1], find, index-step) //	if value less or equal array index value, use binary search for find it
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 8, 11, 13, 14, 19, 27}

	totalStart := time.Now()

	start := time.Now()
	fmt.Println(jumpSearch(arr, 1))
	elapsed := time.Since(start)
	fmt.Println("elapsed time for 1", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 5))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 5", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 8))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 8", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 11))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 11", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 13))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 13", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 19))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 19", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 27))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 27", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 12))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 12", elapsed)

	start = time.Now()
	fmt.Println(jumpSearch(arr, 200))
	elapsed = time.Since(start)
	fmt.Println("elapsed time for 200", elapsed)

	elapsedTotal := time.Since(totalStart)

	fmt.Println("total time elapsed: ", elapsedTotal) //	101.161µs

}
