// pro1 project pro1.go
package main

import (
	"bs"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("rs is %d\n", bs.Bs(5, arr, 0, 7))
}
