package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type ppAt struct {
	Name string
}

func (ppt *ppAt) setName() {
	ppt.Name = "123"
}

//极客时间第15章练习
func main() {
	//structT()
	//unsafeT()
	unsafeT2()
	//yd()
}

//这里我创建一个值的ppt对象,然后验证临时结果和取值等限制
func structT() {
	//这里是无法寻址调用的,因为老衲现在是没有被一个变量指向的,是一个灵魂无处安放的心
	//ppAt{"ppat"}.setName()
	//这里是可以直接调用的 因为老子是有指向者的
	at := ppAt{"ppab"}
	at.setName()
	fmt.Println(at)
}

func mapT() {
	var map1 = map[int]int{1: 1, 2: 2, 3: 3}
	i := map1[1]
	i2 := &i
	//map 值是 1,对于map索引的值来说,是不可以直接取地址的,但是可以赋予一个变量,然后取变量的地址
	//map索引的值 是不可以直接赋值改变的,因为得到的就是一个值呀,除非我对于map的索引值进行赋值操作才可以
	fmt.Println(&i)
	fmt.Println(&i2)
	//fmt.Println(&map1[1]) //此行报错
	//字典字面量或者字典变量的索引值虽然是不可以直接寻址的,但是可以用在自增或者自减的左边
	map1[1]++
	fmt.Println(map1)
}

//这里我练习一下unsafe包里的函数
func unsafeT() {
	//tip1
	at := ppAt{"qws"}
	//打印at对象的name值的指针
	fmt.Println("普通方法取的指针类型", &(at.Name))
	//unsafe包应用
	//只能通过unsafe.Pointer 来进行uintptr的转换
	fmt.Println(&at)
	//结构体的内存地址
	fmt.Println(unsafe.Pointer(&at))
	//uintptr二进制值表示的地址(数值类型)
	atPtr := uintptr(unsafe.Pointer(&at))
	//得到Name字段的起始存储地址----也就是内存地址
	u := atPtr + unsafe.Offsetof((&at).Name)
	//地址转指针类型
	i := (*string)(unsafe.Pointer(u))
	fmt.Println("unsafe取的指针类型", i)

}

func unsafeT2() {
	at := ppAt{"qws"}
	i := &at
	//普通打印%p
	fmt.Printf("普通打印需要p做转换:%p\n", i)
	//查看unsafe.Pointer方法打印的东东
	fmt.Println("打印结构体的指针值", unsafe.Pointer(i))
	atPtr := uintptr(unsafe.Pointer(i))
	fmt.Println(unsafe.Offsetof(i.Name))
	//得到Name字段的起始存储地址----也就是内存地址
	u := atPtr + unsafe.Offsetof(i.Name)
	//地址转指针类型
	h := (*string)(unsafe.Pointer(u))
	//下面简介修改一个结构体属性的三种修改方法,当然还有一种是借助于方法 我们这里暂且不讨论
	//第一种修改方法
	at.Name = "24214"
	//第二种修改方法
	i.Name = "21111"
	//第三种修改方法--已经是Name的指针了,,这特么比发射还屌一点  服气
	*h = "3333"
	fmt.Println(at)
}

//通过反射+unsafe包拿去切片的底层数组的内存地址
func yd() {
	arr := [3]int{1, 2, 3}
	sli := arr[:]
	sliHeader := (*reflect.SliceHeader)(unsafe.Pointer(&sli))
	u := uintptr(unsafe.Pointer(sliHeader.Data))
	fmt.Println(u)
	fmt.Println(sliHeader.Data)
}

//笔记1: 这里讲一下unsafe.Pointer和uintptr的应用
//你看昂 我想要二进制的内存地址 我就需要 uintptr 但是uintptr的值我不能通过指针拿到
// 我需要用unsafe.Pointer做一层包装,这样我就能拿到结构体的二进制值了
//当我有了结构体的指针值,那我就可以直接使用unsafe.Offsetof(指针值.属性),拿到这个属性的二进制偏移量
//结构体属性的二进制偏移量加上结构体的二进制 = 属性相对于起始值的二进制偏移量(因为属性是在结构体里面的)
//现在我要拿属性的指针我就需要 unsafe.Pointer(属性二进制偏移量) 然后强制转换成属性所属的指针类型就可以了
//比如(*string)!

//笔记2:这里讲一下引用指针是否有意义
//以切片为例子  打印 切片值 和打印切片指针的第一个值 得到的都是底层数组的元素地址
// 打印切片指针 得到的是切片结构体的内存地址 这个需要记牢了

//笔记3:这里讲述不可寻址的限制
// 不可寻址的限制 我理解的就是 老子那无处安放的值,但是除了 切片字面量的索引结果的值,
//这个 切片字面量的索引结果的值 虽然是无处安放的心,但是它底层是有数组的,也就是可以直接进行寻址的,不需要赋值变量再寻址
// 对于其他 老子那无处安放的心  直接寻址是不行的 我需要赋值给一个变量 然后取变量的地址

//不可寻址的三大限制:
//  1.不可变的值不可寻址,常量,基本类型的值字面量,字符串(golang 是final)变量的值,函数以及方法的字面量都是如此
//  2.绝大多数的临时结果都是不可寻址的,算术操作的结果,值字面量的表达式都是临时结果,但是有一个例外
// 对切片字面量索引的值 虽然是临时结果 但是可以寻址
//  3.若拿到某值的指针可能破坏程序的一致性,那么就是不安全的,该值不可寻址, 例如字典
// 由于字典的内部机制,对字典的索引结果值的取址操作都是不安全的,另外,获取字面量或标识符代表的函数或者方法的地址也是不安全的
