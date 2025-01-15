// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	for x := 1; x <= 30; x++ {
		for y := 30; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
