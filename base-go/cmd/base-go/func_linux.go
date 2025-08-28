//go:build linux
package main

import (
	"fmt"
)

func platformSpecificFunction() {
	fmt.Println("This is the Linux implementation.")
}
