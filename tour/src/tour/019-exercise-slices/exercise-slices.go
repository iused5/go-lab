package main 

import (
	"fmt"
)

func Pic(dx, dy int) (image [][]uint8) {
	image = make([][]uint8, dy)
	
	for i := range image {
		image[i] = make([]uint8, dx) 
		for j := range image[i] {
			image[i][j] = uint8((i ^ j) * (j ^ i))
		}
		fmt.Printf("%d\n", image[i])
	}
	return
}

func main() {
	fmt.Println(Pic(5, 5))
}

