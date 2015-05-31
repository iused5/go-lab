package main 

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := float64(1)
		prev := float64(0)
	
		fmt.Println("Sqrt ...");
		for math.Abs(prev - z) > 0.00000000001 { 	
			prev = z
			z = z - (((z * z) - x) / (2 * z))
			fmt.Println(z);
		}
		return z, nil	
	}
}

func main() {
	var x float64 = 2
	var s string = fmt.Sprintf("...%v", x);
	fmt.Println(s)
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

