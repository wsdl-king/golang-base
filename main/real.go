package main

import (
	//"flag"
	"fmt"
)

var arg1 string

type hha interface {
	Nam2e()
}
type wai struct {
	Name string
	hha
}
type nei struct {
	age string
}

func (n *nei) Nam2e() {
	n.age = "111"
	//fmt.Println("nei层指针实现的接口", n.age)
}

//接口(hha)所属的结构体的值 (wai),强转成type的时候,如果实现的接口是指针传递,那即使我的wai是值,也会改变的
func main() {
	//wai指针
	var d hha
	d = &nei{"qwe"}

	i := wai{"qws", d}
	i.Nam2e()
	fmt.Println(i.hha.(*nei).age)

	//实现方式1
	//flag.StringVar(&arg1, "arg1", "arg1", "this is arg1")
	//flag.Parse()
	//fmt.Print(arg1)
	//fmt.Println("len",len(os.Args[1:]))
	//fmt.Println("cap",cap(os.Args[1:]))
	//实现方式2 自定义flag对象
	//cmd:=flag.NewFlagSet("",flag.ExitOnError)
	//var i = cmd.String("cmd", "cmd", "this is cmd")
	//cmd.Parse(os.Args[1:])
	//fmt.Println(*i)
	//fmt.Println("len",len(os.Args[1:]))
	//fmt.Println("cap",cap(os.Args[1:]))
	//极客时间第三讲测试--我推荐建立二级目录，这样的话可以go install. test2是 test包下的一个库源文件，但是名字叫test2
	//所以我这里用test2来命令
	//sqrt := test2.Sqrt(1)
	//fmt.Println(sqrt)
	//此处练习 多次引用，还有internal知识点，然后 go install -i 查看可以发生什么!
	//lib2.DD()
	//头条反射
	//reflecct.DD()
	//极客时间重声明
	//_, err = io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//}
	//极客时间值传递
	//var name = *flag.String("name", "everyone", "The greeting object.")
	//flag.Parse()
	//fmt.Println(name)
	//fmt.Println(&name)

	//此设计代码块的知识点
	//对于可重名变量,如果是:=则是可重名变量重新赋值，如果是 = 那就说明是 在原来类型的变量的基础上进行赋值操作
	block := "function"
	{
		//block 称之为 可重名变量。类型不限制
		block := []string{"zero", "one", "two"}
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
	//代码块的知识点结束
	//用点命名 可以省去前缀--极客时间第四章内容知识
	//Print(1)

	//container := map[int]string{0: "zero", 2: "one", 3: "two"}
	fmt.Printf("The element is %q.\n", container[0])

}

//var container = []string{"zer1o", "one", "two"}
var block = "package"
var err error
var container = []string{"zero", "one", "two"}
