package main

import "fmt"

func main() {
	var a, v0, s0, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println(fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	fn := func(t float64) float64 {
		return ((1 / 2) * a * (t * t)) + (v0 * t) + s0
	}

	return fn
}
