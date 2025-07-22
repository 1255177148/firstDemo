package constpkg

import "fmt"

// Gender 枚举实际上就是常量的集合
type Gender string // 给string类型起一个别名，用作枚举的变量类型，这样看着会更清晰
const (
	Boy  Gender = "boy"
	Girl Gender = "girl"
)

func (g *Gender) String() string {
	switch *g {
	case Boy:
		return "boy"
	case Girl:
		return "girl"
	default:
		return "unknown"
	}
}

func (g *Gender) IsBoy() bool {
	return *g == Boy
}

/*
*
测试下枚举
*/
func Demo() {
	g := Boy
	fmt.Println(g.String())
	fmt.Println(g.IsBoy())
}

type Month int

// iota会自动为常量赋值，就是根据该常量在第n行，就赋n-1的值
const (
	Jan Month = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)
