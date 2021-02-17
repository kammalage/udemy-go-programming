package main

import "fmt"

type square struct {
	length int
}

type circle struct {
	radius int
}

type shape interface {
	area() int
}

func main() {
	s := square{
		length: 10,
	}

	c := circle{
		radius: 5,
	}

	info(s)
	info(c)
}

func (s square) area() (area int) {
	area = s.length * s.length
	return
}

func (c circle) area() (area int) {
	area = 3 * c.radius * c.radius
	return
}

func info(s shape) {
	fmt.Println(s.area())
}
