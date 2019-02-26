package main

import "fmt"

//极客时间第14节接口类型的合理运用
type pig struct {
	name     string
	category string
}

type PPet interface {
	//你看 接口里面的方法 仅仅是命名为 接口名字({入参})返回值  这样的格式
	Name() string
	Category() string
	tt() string
}

//这里 无论是传递指针还是传递值 都是可以掉这个方法的,调用这个方法以后,我所进行的操作都是对指针操作
func (pig *pig) SetName(name string) {
	pig.name = name
}

func (pig pig) tt() string {
	return "d"
}

// 这里 是"方法"的定义,带表我的所属者,第二个括号代表我的入参,此例中入参为无参,则视为实现了接口
func (pig pig) Name() string {
	return pig.name
}

func (pig pig) Category() string {
	return pig.category
}

func main() {
	//valuett()
	interExplain()
}

//指针赋值的练习
func point() {
	//声明接口
	var ppet PPet
	//生成一个pig结构体的地址 i就是指针
	i := &pig{"qw", "ew"}
	//指针副本赋值给ppet
	ppet = i
	//调用指针方法
	i.SetName("aaaa")
	//ppet对象和i对象都会发生变化,原因有2个 第一个是 ppet是指针副本,底层还是指向原来的实体
	//第二个原因是SetName是指针方法
	fmt.Print(ppet)
	fmt.Println(i)
}

//值赋值练习
func valuett() {
	//声明接口
	var ppet PPet
	//生成一个pig结构体的值, i就是值
	i := pig{"qw", "ew"}
	//值赋值给ppet
	ppet = i
	//调用指针方法
	i.SetName("aaaa")
	//i的对象会发生变化,因为SetName的入参是指针
	//ppet 接口的值不会发生变化,因为是值传递,传递的是pig对象的副本,pig对象调用指针方法变化关我副本鸟事?
	fmt.Println(ppet)
	fmt.Println(i)

	//前方高能 敲黑板,画重点
	// 虽然ppet被赋予的是pig的值的副本,但是它俩的值和地址也是不同的
	// 地址不同不用我多BB吧
	//值不同,是因为ppet是接口类型的,存储的结构是包含了pig副本的一个结构更加复杂的值,但是ppet包含了dog值的副本
}

//这个函数的意义代表接口的检验
func interExplain() {
	//声明指针类型 但是为空的
	var pig1 *pig
	//赋值给pig2
	pig2 := pig1
	//这里我复制给Pet接口,这个pet它不是为空,只是它特殊的结构里面包含的动态值是空,但是动态类型不是空
	var pet PPet = pig2
	//可以看出 打印的是false
	fmt.Println(pet == nil)
	// 这里 如果
	//  func (pig *pig) tt()string {
	//	 return "d"
	//  }
	//如果是这样的话,我就可以使用含有nil值的接口来调用方法,注意内部是否实例化即可
	//对于静态变量 (这里指代接口声明者pet), 这里pet想调用的方法受制于赋值给它的动态类型,
	//如果赋值给它的动态类型是指针,那么只能调用指针的方法
	//如果赋值给它的动态类型是值, 那么你想用值去调用指针方法是不可能的,因为此时没有实现,所以值只能调用值
	fmt.Println(pet.tt())

}
