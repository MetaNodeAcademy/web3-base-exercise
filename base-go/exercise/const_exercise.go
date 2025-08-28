package exercise

import (
	"fmt"
)

// 方式1
const a int = 1

// 方式2
const b = "test"

// 方式3
const c, d = 2, "Hello"

// 方式4
const e, f bool = true, false

// 方式5
const (
	h    byte = 3
	i         = "Value"
	j, k      = "jj", 4
	l, m      = 5, false
)

const (
	n = 6
)

type AliasInt int

const test = AliasInt(1)

type Gender string

func (g *Gender) isMale() bool {
	return *g == Male
}

func (g *Gender) isFemale() bool {
	return *g == Female
}

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

// iota
type ConnState int
const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

type Month int
const (
	January Month = iota + 1
	February 
	March    
	April    
	May      
	June     
	July     
	August   
	September
	October  
	November 
	December 
)

// 1.6.1 常量与枚举
func DefineConstant() {
	fmt.Println("-----常量与枚举-----")
	var gender Gender = Male
	fmt.Println(gender)
	fmt.Println(gender.isFemale())
	// iota
	fmt.Println("StateNew:", StateNew)
	fmt.Println("StateActive:", StateActive)
	fmt.Println("StateIdle:", StateIdle)
	fmt.Println("StateHijacked:", StateHijacked)
	fmt.Println("StateClosed:", StateClosed)
	fmt.Println("月份:", test)
	fmt.Println("January:", January)
	fmt.Println("February:", February)
	fmt.Println("March:", March)
	fmt.Println("April:", April)
	fmt.Println("May:", May)
	fmt.Println("June:", June)
	fmt.Println("July:", July)
	fmt.Println("August:", August)
	fmt.Println("September:", September)
	fmt.Println("October:", October)
	fmt.Println("November:", November)
	fmt.Println("December:", December)
}
