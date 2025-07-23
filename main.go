package main

import "github.com/1255177148/firstDemo/ClosurePkg"

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

	// 下面是流程控制，if和switch
	//processControl.IfDemo1()
	//processControl.SwitchDemo1()
	//processControl.SwitchDemo2()
	//processControl.SwitchDemo3()
	//processControl.SwitchDemo4()

	// 下面是循环
	//circulatePkg.ForDemo1()
	//circulatePkg.ForDemo2()
	//circulatePkg.ForDemo3()
	//circulatePkg.ForDemo4()
	//circulatePkg.ForDemo5()
	//circulatePkg.ForDemo6()
	//circulatePkg.ForDemo7()

	// 下面是闭包函数
	ClosurePkg.Demo()
}
