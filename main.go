package main

import "fmt"

func circleAround(p int) float64 {
	phi := 3.14
	result := phi * float64(p) * float64(p)
	return result
}

func circleWide(p int)float64{
	phi := 3.14
	result := 2 * phi * float64(p)
	return result
}


func main() {
	fmt.Println(circleAround(10))
	fmt.Println(circleWide(10))
}
