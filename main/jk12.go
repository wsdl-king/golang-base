package main

import (
	"errors"
	"fmt"
)

//函数声明方式1
type f1 func(contents string) int

//函数声明方式2　请比较这两种声明方式的不同
type f2 func(contents string) (n int, e error)

// tip1:函数的签名其实就是函数的参数列表和结果列表的统称,它定义了可用来鉴别不同函数的那些特征,同时也定义了我们与函数的交互方式

//这里声明了一个函数类型
type operation func(x, y int) int

//极客时间第12章节的记录
func main() {

	//highfunction()
	//这个可以看到,我传递的是数组的副本,并不是数组的指针
	//valueT()
	//strings := make([]string, 3, 5)
	//strings2 := make([]string, 3, 6)
	//这里的切片判断相等会报错,这里也说明了,切片函数字典　这些类型是不适用于字典类型的key
	//fmt.Println(strings==strings2)
	//而chan　可以进行比较,说明chan适用于字典类型
	//ints := make(chan int, 1)
	//chanints := make(chan int, 1)
	//fmt.Println(ints == chanints)
	//这里做个笔记,传递的是指针类型,比如　字典,管道,切片　这样的入参在其他的方法中的时候,我们会进行
	//浅拷贝,然后拷贝的其实是指针,底层的实际指向是不变化的
	//下文我们看一个特殊的,传递的是数组(值类型),但是这个数组类型的值是切片类型,哈哈哈,是不是感觉特别有意思
	//我们看结果

	//complexArray1 := [3][]string{
	//	[]string{"d", "e", "f"},
	//	[]string{"g", "h", "i"},
	//	[]string{"j", "k", "l"},
	//}
	//strings := dd(complexArray1)
	//fmt.Println("数组操作后的切片", strings)
	//fmt.Println("操作前的数组", complexArray1)

	//下面展示与一个操蛋的高阶函数写法　哈哈哈哈　贼操蛋
	i := func(x, y int) int {
		return x + y
	}
	o := v3(i)
	fmt.Println(o(1, 2))
	fmt.Println(i(1, 2))
}

//高阶函数1的表达练习
func highfunction() {
	//匿名函数,此函数　入参和出参和operation函数一致,所以就可以视为是一个operation函数
	i := func(x, y int) int {
		return x + y
	}
	//声明
	var d operation
	//赋值(ps　如果你想重声明则需要有另外一个返回值,不然无法重声明)
	d = i
	//v, err := value(1, 2, i)
	//fmt.Println("高阶函数的计算结果为", v)
	//fmt.Println("高阶函数的执行错误为", err)
	//使用
	i2, e := value(1, 2, value2(d))
	fmt.Println("第二次使用高阶函数", i2)
	fmt.Println("第二次使用高阶函数是否返回错误", e)

	//此时我想直接使用value2的函数
	i3 := value2(nil)
	i4 := i3(3, 2)
	fmt.Println("直接使用value2函数返回的结果是", i4)
}

//函数作为入参---高阶函数1
func value(x, y int, op operation) (int, error) {
	if op == nil {
		return 0, errors.New("非法的操作")
	}
	return op(x, y), nil
}

//函数作为返回结果---高阶函数2
func value2(op operation) operation {
	//注意,函数作为入参的时候,你需要if判断是否合法
	if op == nil {
		//不合法我可以定义为一个匿名的内部函数返回!
		return func(d, b int) int {
			return d - b
		}
	}
	//函数作为返回一定要用内部形参,如果用外部的参数的话,会导致返回的函数就是死值函数
	//变量op是外部给的　我可以称之为闭包函数
	return func(d, b int) int {
		return op(d, b)
	}
}

var s [3]int

func valueTrans(a [3]int) [3]int {
	a[0] = 9
	return a
}

//值传递练习查看
func valueT() {
	trans := valueTrans(s)
	fmt.Println("修改之后的值传递", trans)
	fmt.Println("原来的数组的值", s)
}

//这个写法蛮操蛋的哈哈哈哈哈哈哈哈哈哈哈哈
func v3(p operation) operation {
	return p
}

//入参是值类型,但是值类型里面有切片类型
func dd(a [3][]string) [3][]string {
	//这里我修改的是数组内的切片里面的值,如果是这样的话　会影响原来的数组
	//a[0][0]="ll"
	//这里我直接对数组值进行的"替换修改",这里不影响原数组,并且会新生成一个切片和数组
	a[0] = []string{"x", "y", "z"}
	return a
}
