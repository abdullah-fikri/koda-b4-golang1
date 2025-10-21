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
	var p int 
	fmt.Scan(&p)
	// fmt.Println(circleAround(p))
	// fmt.Println(circleWide(p))
	fmt.Println("hasil keliling:", circleAround(p))
	fmt.Println("hasil luas:", circleWide(p))
}
