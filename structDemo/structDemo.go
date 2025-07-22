package structdemo

import "fmt"

//结构体中的字段都有不同的名字，并且字段可以是任意类型，比如结构体本身、指针甚至是函数
type Person struct {
	name   string
	age    int
	slices []any
	p      *int
}

type Student struct {
	name, className, gradeName string
}

// 定义匿名字段，只写字段的类型
type Custom struct {
	int
	string
	Other string // 每个类型只能有一个匿名字段，上面已经有了string类型匿名字段，这个就不能匿名了
}

// 匿名结构体，什么都没定义，是一个空的
var data1 = struct{}{}

// 匿名结构体，不事先定义名字、直接在代码中声明结构体类型
func demo() {
	data := struct {
		Name string
		Age  int
	}{
		Name: "Elvis",
		Age:  18,
	}
	fmt.Println(data)
	fmt.Println(data.Name)
}

// 还有一种匿名结构体，直接写在函数的形参里
func demo2(p struct {
	Name string
	Age  int
}) {
	fmt.Println(p.Name)
}

func demo3() {
	// 这里调用demo2，传入匿名结构体
	demo2(struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 18})
}

type A struct {
	name string
}

type B struct {
	A
	name string
	age  int
}

type C struct {
	A
	B
	age int
}

// 嵌套结构体,可以继承被嵌套的结构体的属性和方法
func Demo4() {
	a := A{name: "a"}
	b := B{name: "b", age: 12, A: a}
	c := C{age: 18, A: a, B: b}
	fmt.Println(c.A.name)
	fmt.Println(c.B.age)
	fmt.Println(c.age)
}

func (a A) String() string {
	return a.name
}

func (a A) SetA() {
	a.name = "v"
}

func (a *A) SetPA() {
	a.name = "aa"
}

func (a *A) GetPA() string {
	return a.name
}

func (b B) StringB() string {
	return b.name
}

func (b *B) SetPBA() {
	b.A.name = "pba"
}

func NewB() B {
	return B{
		age: 18,
		A: A{
			name: "ba",
		},
	}
}

func NewA() A {
	return A{
		name: "a",
	}
}
