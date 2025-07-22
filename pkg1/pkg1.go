package pkg1

import (
	"fmt"
	_"github.com/1255177148/firstDemo/pkg2"
)

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg1 init method invoked")
}

func getPkgName() string {
	fmt.Println("pkg1包的变量初始化了")
	return PkgName
}
