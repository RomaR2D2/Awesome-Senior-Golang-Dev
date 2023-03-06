package oop

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shapes := []Shape{Circle{Radius: 5}, Rectangle{Width: 10, Height: 5}}
	for _, s := range shapes {
		fmt.Println(s.Area())
	}
}

//Здесь мы создаем интерфейс Shape с одним методом Area().
//Затем мы создаем две структуры - Circle и Rectangle - которые реализуют этот интерфейс.
//Круг и прямоугольник имеют разные способы вычисления площади, но они могут быть использованы вместе благодаря интерфейсу Shape.
//Мы создаем срез типа Shape и добавляем в него экземпляры Circle и Rectangle, а затем выводим площадь каждой фигуры.
