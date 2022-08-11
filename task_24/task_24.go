package main

import (
	"fmt"
	"math"
)

/*
24.	Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
We can rewrite the Pythagorean theorem as d=√((x_2-x_1)²+(y_2-y_1)²) to find the distance
between any two points
*/

type Point struct {
	x float64
	y float64
}

/*
	Initialize and return struct
*/
func (p *Point) InitializePoint(x float64, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) InitPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

/*
	Count distance between current point and target point
*/
func (p *Point) Distance(targetPoint Point) float64 {
	var result, l1, l2 float64
	l1 = (targetPoint.x - p.x) * (targetPoint.x - p.x)
	l2 = (targetPoint.y - p.y) * (targetPoint.y - p.y)
	result = math.Sqrt(l1 + l2)
	return result
}

func main() {
	var p1, p2 Point
	p1.InitializePoint(-100.123, 300.234)
	p2.InitializePoint(600.345, 200.456)
	distance_1 := p1.Distance(p2)
	distance_2 := p2.Distance(p1)

	fmt.Printf("P1-P2 distance is: %f\n", distance_1)
	fmt.Printf("P2-P1 distance is: %f\n", distance_2)
}
