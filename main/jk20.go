package main

import (
	"fmt"
	"syscall"
)

//极客时间20章练习--还是关于错误的 我从建造者的角度,去关心 怎样才能给予使用者恰当的错误值的问题
func main() {
	qp()
}

// 构建错误值体系的基本方式有两种, 创建立体的错误类型体系和创建扁平的错误值列表
func qp() {
	//可以去看net.DialTcp的函数实现,返回的是 &OpError 这样的结构体指针,这样的结构体的组成中还包括
	//Err error 的一个属性, 这样的话 我返回一个net.DialTcp的结构体指针,然后结构体指针内还包含链式error
	//net.DialTCP()

	// 用类型建立起树形结构的体系,用统一字段建立起可追溯源的链式错误关联.这是Go语言标准库给予我们的优秀范本
	// 非常有借鉴意义

	// 这种是 Errno 其实是 uintptr 的再定义类型, 每一个常量代表一种系统级别的错误
	// 在zerrors_linux_amd64.go 中有写 syscall包外的代码可以拿到这些错误代表的常量,但是无法改变
	errno := syscall.Errno(1)
	fmt.Print(errno)

}
