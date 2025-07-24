package main

import "fmt"

func main() {
	// 数字类型间可以强转，但高位转低位，高位的数据会被截取掉
	// []byte  []rune  string之间可以无损互转

	// 接口类型转换，也就是Java中的object转别的类型
	var i interface{} = 3
	a, ok := i.(int) // 判断i是否为int并获取i的值
	if ok {
		fmt.Printf("%d is a int", a)
	} else {
		fmt.Println("conversion failed")
	}
	fmt.Println()
	// 或者用switch
	switch s := i.(type) {
	case int:
		fmt.Printf("%d is a int", s)
	case string:
		fmt.Printf("%s is a string", s)
	default:
		fmt.Printf("%d is unknown type", i)
	}

	// 接口转换
	var supplier Supplier = &DigitSupplier{value: "供应商"}
	fmt.Println(supplier.Get())
	b, ok := supplier.(*DigitSupplier)
	fmt.Println(b, ok)

	// 结构体转换，如果结构体的字段相同，可以直接转换值，但是指针不能转换
	personA := PersonA{
		name: "A",
		age:  18,
	}
	fmt.Println(personA)
	personB := PersonB(personA)
	fmt.Println(personB)
}

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value string
}

func (d *DigitSupplier) Get() string {
	return d.value
}

type PersonA struct {
	name string
	age  int
}

type PersonB struct {
	name string
	age  int
}
