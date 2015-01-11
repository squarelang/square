package main

import "fmt"

func add(i int, j int) int {
	return i + j
}

var varName float64

type geometry interface {
	area() float64
	scale(float64) float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) scale(varName float64) float64 {
	s.width = s.width * varName
	s.height = s.height * varName
	return s
}

func (c circle) area() float64 {
	return c.width * c.height
}

func (c circle) scale(varName float64) float64 {
	c.height = c.height * varName
	return c
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

var mysquare = square{width: 1, height: 2}

var mysquare = square{
	width: struct{
		hello: 1,
		goodbye: 2},
	height: 2
}

var mysquare = square{
	width: 1,
	height: 2
}
