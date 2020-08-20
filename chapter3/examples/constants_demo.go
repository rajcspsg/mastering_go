package examples

import (
	"fmt"
)

type Digit int
type Power2 int
type Power3 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)


const (
	Zero Digit = iota
	One
	Two
	Three
	Four
)

const (
	P2_0 Power2 = 1 << iota
	_
	P2_2
	_
	P2_4
	_
	P2_6
)

const (
	P3_0 Power3 = 3 * iota
	P3_1
	P3_2
)

func ConstantsDemo() {
	fmt.Println(One)
	fmt.Println(Two)

	fmt.Println("2^0:", P2_0)
	fmt.Println("2^2:", P2_2)
	fmt.Println("2^4:", P2_4)
	fmt.Println("2^6:", P2_6)


	fmt.Println("3^0: ", P3_0)
	fmt.Println("3^1: ", P3_1)
	fmt.Println("3^2: ", P3_2)

}
