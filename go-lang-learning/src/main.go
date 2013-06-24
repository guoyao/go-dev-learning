/**
 * User: guoyao
 * Time: 06-20-2013 12:16
 * Blog: http://www.guoyao.me
 */

package main

import (
	"fmt"
	"math"
	"concurrency"
)

type Shape interface {
	perimeter() float64
	area() float64
}

type Circle struct {
	x,  y,  r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type MultiShape struct {
	shapes []Shape
}

type Person struct {
	name string
}

type Android struct {
	Person
	model string
}

func main() {
//	test()
//	testSlice()
//	testMap()
//	fmt.Println(fibonacci(40))
//	fmt.Println(add([]int{1, 2, 3, 4, 5}...))
//	closure()

//	nextEven := makeEvenGenerator()
//	fmt.Println(nextEven()) // 0
//	fmt.Println(nextEven()) // 2
//	fmt.Println(nextEven()) // 4

//	fmt.Println(factorial(5))

//	x, y := 0, 1
//	fmt.Println(x, y)
//	swap(&x, &y)
//	fmt.Println(x, y)

//	c := Circle{0, 0, 5}
//	fmt.Printf("%f - %f\n", c.perimeter(), c.area())
//
//	r := Rectangle{0, 0, 5, 5}
//	fmt.Printf("%f - %f\n", r.perimeter(), r.area())
//
//	fmt.Println(totalAreas(&c, &r))
//
//	m := MultiShape{[]Shape{&c, &r}}
//	fmt.Printf("%f - %f\n", m.perimeter(), m.area())
//
//	m2 := MultiShape{[]Shape{&m, &c}}
//	fmt.Printf("%f - %f\n", m2.perimeter(), m2.area())

//	p := Person{"Guoyao"}
//	android := new(Android)
//	android.talk()

	concurrency.Run()
}

func test() {
	x := [5]float64{1, 2, 3, 4, 5}
	result := float64(1)
	for i := 0; i < len(x); i++ {
		result *= x[i]
	}
	fmt.Println(result)
	fmt.Println(result/float64(len(x)))

	result = 0
	for _, value := range x {
		result += value
	}
	fmt.Println(result)
}

func testSlice() {
	slice1 := []int {1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	slice3 := make([]int,3, 9)
	fmt.Println(len(slice3))
}

func testMap() {
	x := make(map[string]int)
	x["key1"] = 1
	fmt.Println(x["1"])
	y := make(map[string]string)
	y["guoyao"] = "Guoyao Wu"
	y["lingxuan"] = "Lingxuan Wu"
	if name, ok := y["guoyao"]; ok {
		fmt.Printf("\"%s\" exsits\n", name)
	}
}

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}
	return fibonacci(n - 2) + fibonacci(n - 1)
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func closure() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(5, 10))
}

func makeEvenGenerator() func () uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x*factorial(x - 1)
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c *Circle) float64 {
	return math.Pi*c.r*c.r
}

func (c *Circle) area() float64 {
	return math.Pi*c.r*c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, shape := range m.shapes {
		perimeter += shape.perimeter()
	}
	return perimeter;
}

func totalAreas(shapes ...Shape) float64 {
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.area()
	}
	return totalArea
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}
