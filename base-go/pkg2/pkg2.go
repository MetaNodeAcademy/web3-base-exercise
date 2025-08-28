package pkg2

import "fmt"

const PkgName string = "pkg2"

var PkgNameVar string = GetPkgName()

func init() {
	fmt.Println("pkg2 init function has been called")
}

func GetPkgName() string {
	fmt.Println("pkg2.getPkgName has been called")
	return PkgName
}
