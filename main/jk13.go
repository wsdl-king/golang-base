package main

import "fmt"

//极客时间36讲,此节内容讲述的是"结构体"
type Qw func(x, y int) int

//声明一个带函数的
type ss struct {
	Name string
}
type qwe struct {
	Age int
	ss
}

type cat struct {
	Name string
}
type animal struct {
	Name string
	cat
	dog
}
type dog struct {
	Name string
}

func main() {
	//i2 := ss{"99"}
	//i2.Add()
	//此处i2是值类型,但是对i2结构体内部的name做变化也是可以改变值的
	//注意注意.只有在入参的时候才有　值或者指针一说,我自己声明出来的结构体不当入参进行调用,
	//那么　无论我自己的结构体我做什么,还是会影响到我的内部属性的
	//i2.Name = "1313131"
	//fmt.Println(i2)

	//这里我做嵌入字段的打印
	//innerPrint()
	//重名的结构体嵌入打印
	//innerStruct()
	TwoStatement()
}

func TwoStatement() {
	//两种声明方式--都是指针
	//地址调用,但是也是值传递,值传递依赖于Add的方法的入参,看是指针还是值
	var s = &ss{}
	fmt.Println(s)
}

func (s ss) Add() {
	s.Name = "123"
}

//重写了此结构体的String方法
//重写以后,打印就是用的我自己写的方法--注意这里只能是隐式调用,如果我的变量声明为值类型的话,此次调用是没用的
//我的变量声明为指针类型的话　这里的隐式调用是成功的
// 还有 如果我的String是值类型的话,无论 ss是什么类型,都是可以调用成功的
// 即 结构体的指针 无视结构体的 值/指针方法. 结构体的值 只能适应值方法
func (s ss) String() string {
	return "d"
}

// qwe持有 ss  如果qws没有重写String方法,则会调用内联结构体的String方法,但是此时我重写了qwe的String方法
//所以 下次的打印会打印我qwe重写的String方法
func (s *qwe) String() string {
	return "qwe"
}

//嵌入字段的打印
//　笔记2:这里可以看到,我没有重写qwe结构体的String方法,只重写了s结构体的String方法,但是
//如果qws结构体中引用的是s的指针的话,那么String方法一并被替换,因为他们的公共接口都是
//                 type Stringer interface {
//	                       String() string
//                 }
func innerPrint() {
	var qw qwe
	var sse ss
	//qw.ss.Name = "123"
	qw.Age = 1
	i := &qwe{9, sse}
	//这里打印的qw,因为qw是值类型,并且内部也没有其他结构体的指针引用,String方法对他根本没有类型
	fmt.Println(qw)
	//这里打印的i 因为i是指针类型,所以我打印的是重写以后print
	fmt.Println(i)
}

//注意注意 看这个 innerStruct函数和ani函数,并且观察 animal cat dog的类型
// 可以得出结论  外层实体的内联实体是指针类型时,调用外层实体的值传递方法,在值传递方法内对指针操作
//照样是可以改变内联实体的数据的,但是如果外层实体的内联实体不是指针类型的时候,就不会改变了
//想要改变的话,就需要把外联实体的方法改成指针传递!
func innerStruct() {
	var an animal
	an.Name = "an"
	//an.cat=&cat{"ca111"}
	//an.cat.Name = "cat"
	//an.dog.Name = "dog"
	ani(an)
	fmt.Println(an)
	//fmt.Println(*an.cat)
}
func (a animal) ani() {
	a.Name = "a1"
	//a.dog.Name = "dog1"
	a.cat.Name = "cat1"
}
func ani(a animal) {
	a.Name = "a1"
	//a.dog.Name = "dog1"
	a.cat.Name = "cat1"
}
