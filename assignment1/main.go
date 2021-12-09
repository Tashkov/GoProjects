package main

import "fmt"

type shape interface {
	getArea() (float64, error)
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	tr := triangle{
		height: 10,
		base:   10,
	}

	sq := square{
		sideLength: 10,
	}

	printArea(tr)
	printArea(sq)

}

func printArea(s shape) (int, error) {
	result, err := s.getArea()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return fmt.Println(result)
}

func (tr triangle) getArea() (float64, error) {
	return float64(0.5 * tr.base * tr.height), nil
}

func (sq square) getArea() (float64, error) {
	return float64(sq.sideLength * sq.sideLength), nil
}
