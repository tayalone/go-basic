package gotypes

import "fmt"

/*
DemoArray show how to use Array in Go
*/
func DemoArray() {
	/* array in Go is IMMUTABLE */

	/* declare empty int[] */
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Printf("arr1 type :%T, value: %v\n", arr1, arr1)

	/* declare define int[] */
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr2 type :%T, value: %v\n", arr2, arr2)
}
