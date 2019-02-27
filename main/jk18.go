package main

import "fmt"

//极客时间第18章练习
func main() {
	//szT()
	//qpT()
	//qppj()
	//switchT1()
	//switchT2()
	switchT5()
}

// 这里测试rang值类型demo--数组
func szT() {
	ints := [3]int{1, 2, 3}
	size := len(ints) - 1
	// i 是索引, e是值
	//这里的 i e 是 ints后的副本值,因为ints是值对象,然后golang是值传递,所以是副本值
	for i, e := range ints {
		//也就是说i,e在每一行的值已经确定了
		if i == size {
			ints[0] += e
		} else {
			ints[i+1] += e
		}
	}
	fmt.Println(ints)

}

// 这里测试rang值类型demo--切片
func qpT() {
	ints := []int{1, 2, 3}
	size := len(ints) - 1
	// i 是索引, e是值
	for i, e := range ints {
		// 这里的e的值是动态的 与其说是动态的这样解释还不如这样解释
		// range 右边的表达式 是指针 求的值复制给 i,e他们都是指针副本 底层指向相同的数组元素
		if i == size {
			ints[0] += e
		} else {
			ints[i+1] += e
		}
	}
	fmt.Println(ints)
}

// 切片拼接高级操作
func qppj() {
	ints := []int{1, 2, 3}
	// i 是索引, e是值,这里演示了一个切片用range进行拼接的操作
	for i := range ints {
		ints = append(ints, i)
	}
	fmt.Println(ints)
}

//笔记: for-range
// tip1: 对于有索引项的数据表达式进行rang操作的时候,生成的是两个值,i,e  左边的值代表索引,右边的值代表具体的值
// tip2: range表达式只会在for语句开始执行的时候被求值一次,无论后面有多少迭代,当然这里要注意赋值的时候是值传递还是指针传递
// tip3: range表达式的结果会被复制,也就是说,被迭代的对象是range表达式结果的副本而不是原值

// switch 函数的练习
// 这里的byte和 uint8本质类型是一个类型,所以不可以取值
//func switchT1() {
//	value6 := interface{}(byte(127))
//	switch t := value6.(type) {
//	case uint8, uint16:
//		fmt.Println("uint8 or uint16")
//	case byte:
//		fmt.Printf("byte")
//	default:
//		fmt.Printf("unsupported type: %T", t)
//	}
//}

//此例子可以绕过 case 值重复的编译检查,因为我的case值 是表达式 而不是字面量
// 但是 处于最上面的表达式会被打印值 --- case a,b,c 意为多个值满足其中一个就可以
func switchT2() {
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}
}

//此例子不可以绕过case值重复的编译检查,因为我的case值 是字面量 而不是表达式
//下面是无法通过编译的
//func switchT3() {
//	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
//	switch value5[4] {
//	case 1, 2, 3:
//		fmt.Println("0 or 1 or 2")
//	case 4, 5, 6:
//		fmt.Println("2 or 3 or 4")
//	case 1, 2, 9:
//		fmt.Println("4 or 5 or 6")
//	}
//}

//下面我们着重来讲一下 switch 和case之间的类型转换
// tip1: 这样是无法通过编译的 1+3会被自动转换成int类型,而case 是int8的类型,所以呢 所以呢不能判断相等!
//当然 除非你强转
func switchT4() {
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch (int8)(1 + 3) {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}
}

// tip2: 这里展示了case的隐式转换,我们看一下例子
// case 的表达式是没有类型的 而 switch 的表达式是有类型的, 所以case 的表达式会转换成switch表达式的类型
func switchT5() {
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// 这里我是int8类型的
	switch value5[0] {
	case 0, 1, 2:
		fmt.Println("0 or 1 or 2")
	case 3, 4, 5:
		fmt.Println("2 or 3 or 4")
	case 6, 7, 8:
		fmt.Println("4 or 5 or 6")
	}
}
