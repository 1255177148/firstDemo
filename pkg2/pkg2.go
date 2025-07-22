package pkg2

import "fmt"

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg2 init method invoked")
}

func getPkgName() string {
	fmt.Println("pkg2包的变量初始化了")
	return PkgName
}

