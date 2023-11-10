package main

import (
	"fmt"
	"math"
)

type cone struct {
	radius, height float64
}

func (c *cone) area() float64 {
	A := (2 * c.radius * c.height * math.Pi) +( 2 * c.radius * c.radius * math.Pi)
	return A
}

func (c cone) volume() float64 {
	V := (c.radius * c.radius * c.height * math.Pi)
	return V
}

func main() {
	c := cone{radius: 10.0, height: 5.0}

	fmt.Println("area: ", c.area())
	fmt.Println("volume:", c.volume())

	cpoint := &c
	fmt.Println("area: ", cpoint.area())
	fmt.Println("volume:", cpoint.volume())
}