package circulatePkg

import "fmt"

/*
普通的for循环
*/
func ForDemo1() {
	for i := 0; i < 10; i++ {
		fmt.Println("for循环方式一，第", i, "次循环")
	}
}

/*
将变量改变写在循环内的for循环
*/
func ForDemo2() {
	b := 1
	for b < 10 {
		fmt.Println("for循环方式二，第", b, "次循环")
		b++
	}
}

/*
遍历数组
*/
func ForDemo3() {
	var a [10]string
	a[0] = "a"
	for i := range a {
		fmt.Println("当前下标", i)
	}
	for i, e := range a {
		fmt.Println("a[", i, "]=", e)
	}
}

/*
遍历切片
*/
func ForDemo4() {
	s := make([]string, 10)
	s[0] = "s"
	for i := range s {
		fmt.Println("当前下标", i)
	}
	for i, e := range s {
		fmt.Println("s[", i, "]=", e)
	}
}

/*
遍历map
*/
func ForDemo5() {
	m := make(map[string]string)
	m["a"] = "Hello,a"
	m["b"] = "Hello,b"
	m["c"] = "Hello,c"
	for k := range m {
		fmt.Println("当前key:", k)
	}
	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}
}

/*
使用break
*/
func ForDemo6() {
	for i := 0; i < 10; i++ {
		fmt.Printf("不使用标记，外层循环，i=%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("不使用标记，内层循环，j=%d\n", j)
			break
		}
	}

loop:
	for i := 0; i < 10; i++ {
		fmt.Printf("使用标记，外层循环i=%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("使用标记，内层循环j=%d\n", j)
			break loop
		}
	}
}

/*
使用continue
*/
func ForDemo7() {
	for i := 0; i < 10; i++ {
		fmt.Printf("不使用标记，外层循环，i=%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("不使用标记，内层循环，j=%d\n", j)
			if j == 3 {
				continue
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

loop:
	for i := 0; i < 10; i++ {
		fmt.Printf("使用标记，外层循环i=%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("使用标记，内层循环j=%d\n", j)
			if j == 3 {
				continue loop
			}
			fmt.Println("使用标记，内部循环，在continue之后执行")
		}
	}
}
