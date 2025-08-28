//go:build windows
package main

import (
	"fmt"
)

func platformSpecificFunction() {
	fmt.Println("This is the Windows implementation.")
}
