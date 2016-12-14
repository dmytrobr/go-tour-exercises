//https://tour.golang.org/flowcontrol/8


package main


import (
    "fmt"
    "math"
)


func Sqrt(x float64) (float64, int) {
    var startingPoint = x/2
    result := float64(startingPoint)
    var delta = 1.0
    var i=0
    for i=0;delta > 0.00001 || delta < -0.00001; i++ {
   	 previousResult := result
   	 delta=((previousResult*previousResult)-x)/(2*previousResult)
   	 result = previousResult - delta
   	 fmt.Printf("result: %v, previousResult: %v, delta: %v\n", result, previousResult, delta)
    }
    return result, i


}


func main() {
    fmt.Println(Sqrt(15))
    fmt.Println(math.Sqrt(15))
}
