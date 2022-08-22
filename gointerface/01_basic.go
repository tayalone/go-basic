package gointerface

import (
	"fmt"
	"strconv"
)

/*
DemoInterface1 Showe How Go Interface Work
*/
func DemoInterface1() {
	//// ----------- empty interface -------------------
	var itf interface{}

	itf = 1
	fmt.Printf("type: %t, value %v\n", itf, itf)

	itf = "1"
	fmt.Printf("type: %t, value %v\n", itf, itf)

	itf = true
	fmt.Printf("type: %t, value %v\n", itf, itf)

	itf = make([]int, 3, 6)
	fmt.Printf("type: %t, value %v\n", itf, itf)
	//// -----------------------------------------------
}

type Stringer interface {
	String() string
}

type S string

func (s S) String() string {
	return string(s)
}

type I int

func (i I) String() string {
	return strconv.Itoa(int(i))
}

/*
DemoInterface2 Showe How Go Interface Work
*/
func DemoInterface2() {
	//// ---------- Use Interface for Design Type's Bahavior -----
	var s S = "1"

	var i I = 1

	stringerArr := []Stringer{s, i}

	for i, v := range stringerArr {
		fmt.Println("Stringer no:", i+1, "value is:", v.String())
	}

	////---------------------------------------------------------
}
