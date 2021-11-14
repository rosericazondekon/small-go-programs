package main

import (
	"fmt"
)

func main() {
	var acc, vnull, snull, time float64
	fmt.Printf("Enter value for acceleration: ")
	fmt.Scanf("%f\n", &acc)
	fmt.Printf("Enter value for initial velocity: ")
	fmt.Scanf("%f\n", &vnull)
	fmt.Printf("Enter value for initial displacement: ")
	fmt.Scanf("%f\n", &snull)
	fmt.Printf("Enter value for time in seconds: ")
	fmt.Scanf("%f\n", &time)

	fn := GenDisplaceFn(acc, vnull, snull)

	fmt.Printf("The displacement is: %f\n", fn(time))
}

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	myfn := func(t float64) float64 {
		s := 0.5*a*t*t + v_o*t + s_o
		return s
	}
	return myfn
}
