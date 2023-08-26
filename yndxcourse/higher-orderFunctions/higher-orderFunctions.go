package main

import (
	"fmt"
	"math"
)

const dimension = 10

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(figure figures) (functor func(float64) float64, isKnown bool) {
	isKnown = true
	switch figure {
	case square:
		functor = func(dimension float64) float64 {
			return dimension * dimension
		}
	case circle:
		functor = func(radius float64) float64 {
			return math.Pi * radius * radius
		}
	case triangle:
		functor = func(side float64) float64 {
			return math.Sqrt(3) * side * side / 4
		}
	default:
		functor = nil
		isKnown = false
	}
	return
}

func main() {
	myFigures := []figures{square, circle, triangle, figures(10)}
	for _, myFigure := range myFigures {
		areaFunc, ok := area(myFigure)
		if !ok {
			fmt.Println("Ошибка")
		} else {
			myArea := areaFunc(dimension)
			fmt.Println(myArea)
		}
	}
}
