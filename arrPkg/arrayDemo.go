package arrPkg

// 声明数组
var arr1 [5]int

// 声明数组并初始化
var arr2 = [2]int{1, 2}

// 声明数组并初始化，用...代替长度，编译器会自动根据元素个数推断长度
var arr3 = [...]int{1, 2, 3}
