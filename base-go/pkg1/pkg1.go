package pkg1

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

func init() {
	//fmt.Println("pkg1 init function has been called")
}

func getPkgName() string {
	//fmt.Println("pkg1.PkgNameVar has been called")
	//fmt.Println("pkg2.getPkgName -> " + pkg2.GetPkgName())
	return PkgName
}
