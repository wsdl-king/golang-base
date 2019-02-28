package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

//极客时间19章练习,19章讲述的主要是 错误处理,错误处理分为两章先学第一章
// 使用者的角度 怎样处理好错误值
func main() {
	//jkT()
	//这里返回一个新的error值,内部其实是sprintf包装了一下String 然后errors.New直接生成了
	//jktest3()
	sl2()
}

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

// func 这里复习一下方法和函数的区别写法
func diff() {
	//error是一个接口
	// 这里调用的方式为: 接口.函数(实现此接口的结构体)--指针哦 如果是值是没有实现此接口的
	error.Error(errors.New("错误了"))
	//实现方式2,我直接把这个接口的实现给搞出来,然后 实现.方法
	//具体的结构体实现
	qw := errors.New("23")
	//实现.方法
	e := qw.Error()
	fmt.Println(e)
}

//极客时间的练习例子
// 在我们的内部 对于error的判断通常使用 nil 来进行判断 然后处理我们相应的逻辑
func jkT() {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}

func underlyingError(err error) error {
	switch err := err.(type) {
	//指针类型
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

// Errno 代表某种错误的类型。
type Errno int

func (e Errno) Error() string {
	return "errno " + strconv.Itoa(int(e))
}

func jktest3() {
	var err error
	// 示例1。
	_, err = exec.LookPath(os.DevNull)
	fmt.Printf("error: %s\n", err)
	//接口.(*实体)
	//这特么是个类型转换啊啊啊啊啊啊啊 我曹
	// interface{}(value).(Type) 在这个Type里面我是不会检查type的值的
	//接口.(Type) 进行类型转换 返回两个值 我会进行判断的
	//接口转Type
	if execErr, ok := err.(*exec.Error); ok {
		execErr.Name = os.TempDir()
		execErr.Err = os.ErrNotExist
	}
	fmt.Printf("error: %s\n", err)
	fmt.Println()

}

//示例2
func sl2() {
	// 示例2。
	var err error
	err = os.ErrPermission
	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
	os.ErrPermission = os.ErrExist
	// 上面这行代码修改了os包中已定义的错误值。
	// 这样做会导致下面判断的结果不正确。
	// 并且，这会影响到当前Go程序中所有的此类判断。
	// 所以，一定要避免这样做！
	if os.IsPermission(err) {
		fmt.Printf("error(permission): %s\n", err)
	} else {
		fmt.Printf("error(other): %s\n", err)
	}
}

//常量Error的练习
func consError() {
	// 示例3。
	const (
		ERR0 = Errno(0)
		ERR1 = Errno(1)
		ERR2 = Errno(2)
	)
	var myErr error = Errno(0)
	switch myErr {
	case ERR0:
		fmt.Println("ERR0")
	case ERR1:
		fmt.Println("ERR1")
	case ERR2:
		fmt.Println("ERR2")
	}

	//笔记   怎样判断一个错误值代表的是哪一类错误
	// tip1: 对于类型在已知范围内的一系列错误值,一般使用类型断言表达式或类型switch语句来判断
	// tip2: 对于已有相应变量且类型相同的一系列错误值,一般直接使用判等操作来判断 switch case 值相等
	// tip3: 对于没有相应变量且类型未知的一系列错误值,只能使用其错误信息的字符串表示形式做判断

}
