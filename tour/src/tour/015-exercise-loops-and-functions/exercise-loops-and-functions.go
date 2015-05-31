package main 

import (
	"fmt"
	"math"
)

func SqrtTemp(x float64) float64 {
	z := float64(1)
	
	fmt.Println("Sqrt Temp ...");
	for i := 0; i < 10; i++ {
		z = z - (((z * z) - x) / (2 * z))
		fmt.Println(z);
	}
	
	return z
}

func Sqrt(x float64) float64 {
	z := float64(1)
	prev := float64(0)

	fmt.Println("Sqrt ...");
	for math.Abs(prev - z) > 0.00000000001 { 	
		prev = z
		z = z - (((z * z) - x) / (2 * z))
		fmt.Println(z);
	}
	
	return z
}	

func main() {
	fmt.Println(Sqrt(10))
	fmt.Println(SqrtTemp(10))
}

