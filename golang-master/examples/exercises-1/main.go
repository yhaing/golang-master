package main

import (
	"fmt"
)

func main() {
	a := []string{"I", "am", "stupid", "and", "weak"}
	for k, v := range a {
		if v == "stupid" {
			a[k] = "smart"
		}
		if v == "weak" {
			a[k] = "strong"
		}
	}
	fmt.Println(a)
}
