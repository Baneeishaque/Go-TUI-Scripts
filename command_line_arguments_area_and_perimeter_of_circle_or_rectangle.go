package main

import "fmt"
import "os"
import "strconv"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius

}

func (re Rectangle) Area() float64 {
	return re.Length * re.Width
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius

}

func (re Rectangle) Perimeter() float64 {
	return 2 * (re.Length + re.Width)

}

func main() {
	if os.Args[1] == "circle" {
		ra, _ := strconv.ParseFloat(os.Args[2], 64)
		cir := Circle{ra}
		fmt.Printf("Area is %v\n", cir.Area())
		fmt.Printf("Perimeter is %v", cir.Perimeter())
	} else {
		if os.Args[1] == "rectangle" {
			le, _ := strconv.ParseFloat(os.Args[2], 64)
			wi, _ := strconv.ParseFloat(os.Args[3], 64)
			re := Rectangle{le, wi}
			fmt.Printf("Area is %v\n", re.Area())
			fmt.Printf("Perimeter is %v", re.Perimeter())
		}
	}

	//	fmt.Printf("Perimeter is %v", Perimeter)

}
