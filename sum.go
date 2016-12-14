//Find sum of elements to the right and left
package main

import "fmt"

func doubleSide(a []int, result []int, sum int, i int) {
	size := len(a)
	var tmpSumFwd int
	var tmpSumBackward int
	for i := 0; i < size-1; i++ {

		backwardIndex := size - i - 1
		tmpSumFwd = tmpSumFwd + a[i]
		tmpSumBackward = tmpSumBackward + a[backwardIndex]

		result[i+1] = result[i+1] + tmpSumFwd
		result[backwardIndex-1] = result[backwardIndex-1] + tmpSumBackward

		fmt.Println(result)
	}
}
func main() {
	a := []int{3, 1, 2, 12, 1, 1, 1, 12, 3}
	size := len(a)
	result := make([]int, size)

	doubleSide(a, result, 0, 0)

	fmt.Printf("For loop result: %v \n", result)

	result = make([]int, size)

	_ = sumArray(a, result, 0, 0)

	fmt.Printf("Recursion result: %v", result)

}
func sumArray(a []int, result []int, sum int, i int) int {

	sum = sum + a[i]
	if i < len(a)-1 {
		sum = sumArray(a, result, sum, i+1)
	}
	result[i] = sum - a[i]
	return sum
}
