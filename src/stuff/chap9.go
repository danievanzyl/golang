package main

import (
	"fmt"
	"math"
	"time"
)


type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{1,1,4}
	d := Circle{5,6,10}


	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())
	fmt.Println(d.x, d.y, d.r)
	fmt.Println(d.area())
	fmt.Println(time.Now())
}