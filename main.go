package main

import (
	"github.com/1255177148/firstDemo/processControl"
)

func main() {
	//var a complex128 = 1.01 + 0.1i
	//b := 2 + 0.2i
	//c := complex(3, 0.5)
	//x := real(c)
	//y := imag(c)
	//fmt.Println(a, b, c, x, y)
	//
	//var s = "语言"
	//var runes = []rune(s)
	//fmt.Println(runes)
	//fmt.Println(len(runes))
	//var bytes []byte = []byte(s)
	//fmt.Println(bytes)
	//fmt.Println(len(bytes))
	//fmt.Println(string(bytes))

	//demoPoint() // 指针操作

	// 下面是结构体方法操作
	//structdemo.Demo4()
	//bb := structdemo.NewB()
	//bb.SetA()
	//fmt.Println(bb.A.String())
	//bb.SetPA()
	//fmt.Println(bb.A.String())
	//bb.SetPBA()
	//fmt.Println(bb.A.String())
	//aa := structdemo.NewA()
	//pa := &aa
	//fmt.Println(aa.String())
	//fmt.Println(pa.String())
	//pa.SetA()
	//fmt.Println(aa.String())
	//fmt.Println(pa.String())
	//pa.SetPA()
	//fmt.Println(aa.String())
	//fmt.Println(pa.String())
	processControl.IfDemo1()
	processControl.SwitchDemo1()
	processControl.SwitchDemo2()
	processControl.SwitchDemo3()
	processControl.SwitchDemo4()
}
