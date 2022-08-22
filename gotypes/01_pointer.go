package gotypes

import "fmt"

func passByPntr(p *string) {
	*p = *p + " passByPntr"
}

func passByValue(s string) string {
	return s + " passByValue"
}

/*
DemoPointer show how pointer work in go
*/
func DemoPointer() {
	s1 := "my string"
	p1 := &s1

	fmt.Println("-------------- test pntr ----------------------------")
	fmt.Printf("s1 type: %T, value: %s\n", s1, s1)
	fmt.Printf("p1 type: %T, value: %x, ref value %s\n", p1, p1, *p1)
	fmt.Println("-----------------------------------------------------")

	fmt.Println("-------------- test passByValue ----------------------------")
	s2 := passByValue(s1)
	fmt.Println("s2 value is: ", s2)
	fmt.Printf("s1 address: %x, s2 address %x, is equal address %v\n", &s1, &s2, &s1 == &s2)
	fmt.Println("-----------------------------------------------------")

	fmt.Println("-------------- test passByPointer ----------------------------")
	passByPntr(p1)
	fmt.Printf("p1 ref value %s\n", *p1)
	fmt.Printf("s1 value %s\n", s1)
	fmt.Printf("is s1 eq *p1: %v\n", s1 == *p1)
	fmt.Println("-----------------------------------------------------")
}
