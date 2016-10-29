package main

import (
	"fmt"
	"math"
)

type AbsInterface interface {
	Abs() float64
}

type Point struct {
	x float64
	y float64
}

type PointName struct {
	Point
	name string
}

type Point3D struct {
	x float64
	y float64
	z float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *PointName) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point3D) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	p1 := Point{x: 1, y: 1}
	p2 := PointName{Point{2, 2}, "twotwo"}
	var p3 = Point3D{x: 1, y: 2, z: 3}

	fmt.Println("p1", p1.Abs())
	fmt.Println("p2", p2.Abs())
	fmt.Println("p3", p3.Abs())

	var ai AbsInterface
	p4 := new(Point)
	ai = p4
	fmt.Println("ai aka p4", ai)

	p5 := Point{11, 11}
	fmt.Println("p5", p5)
	ai = &p5
	fmt.Println("ai aka p5", ai, ai.Abs())
}
