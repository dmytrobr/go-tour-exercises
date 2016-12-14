//https://tour.golang.org/methods/20
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, int, error) {
	if x < 0 {

		return 0, 0, ErrNegativeSqrt(x)
	}

	var startingPoint = x / 2
	result := float64(startingPoint)
	var delta = 1.0
	var i = 0
	for i = 0; delta > 0.001 || delta < -0.001; i++ {
		previousResult := result
		delta = ((previousResult * previousResult) - x) / (2 * previousResult)
		result = previousResult - delta
	}
	return result, i, nil

}

func main() {

	result, i, err := Sqrt(-2)
	fmt.Println(Sqrt(2))

	fmt.Println(err)
	fmt.Println(result)
	fmt.Println(i)

}
