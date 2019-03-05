package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//极客时间45,OS系统相关的包
// 问题 可应用于File值的操作模式有那些?
// os.O_RDONLY,os.O_WRONLY,os.O_RDWR,os.TRUNC,os.append,os.create,os.sync,os.excl(与create一起使用给定路径不能存在文件),
//只读 只写 读写 文件存在就清空 追加 创建 打开文件同步IO

func main() {
	jk89()

}
func simpleOST() {

	//os.Create的实现...//return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	//os.Create()
	//读写模式 创建截断  |是位操作符, 权限666
	//os.Open的实现... // return OpenFile(name, O_RDONLY, 0)
	//只读模式打开 权限是0
	//os.Open()
	//os.OpenFile 函数 的第三个参数表示权限
	//type FileMode uint32  有一个这样的数据类型 低9位都是权限类型的设置
	//类似于0666 0777这样的标示方法都是8进制
	// 为例: 低9位中,分为三组 高到低  每组依次代表了 读,写,执行.如果相应的是1 就代表权限开启
	// 而三个组 代表的用户依次是 文件所有者,文件所有者所属的用户组,其他用户对该文件的访问权限

}

//极客时间88练习
func jk88() {
	fileName1 := "something2.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	fmt.Printf("The file path: %s\n", filePath1)
	fmt.Println()

	// 示例1。
	contents0 := "OpenFile is the generalized open call."
	flagDescList := []flagDesc{
		{
			os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
			"os.O_WRONLY|os.O_CREATE|os.O_TRUNC",
		},
		{
			os.O_WRONLY,
			"os.O_WRONLY",
		},
		{
			os.O_WRONLY | os.O_APPEND,
			"os.O_WRONLY|os.O_APPEND",
		},
	}

	for i, v := range flagDescList {
		fmt.Printf("Open the file with flag %s ...\n", v.desc)
		file1a, err := os.OpenFile(filePath1, v.flag, 0666)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The file descriptor: %d\n", file1a.Fd())

		contents1 := fmt.Sprintf("[%d]: %s ", i+1, contents0)
		fmt.Printf("Write %q to the file ...\n", contents1)
		n, err := file1a.WriteString(contents1)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The number of bytes written is %d.\n", n)

		file1b, err := os.Open(filePath1)
		fmt.Println("Read bytes from the file ...")
		bytes, err := ioutil.ReadAll(file1b)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("Read(%d): %q\n", len(bytes), bytes)
		fmt.Println()
	}

	// 示例2。
	fmt.Println("Try to create an existing file with flag os.O_TRUNC ...")
	file2, err := os.OpenFile(filePath1, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("The file descriptor: %d\n", file2.Fd())

	fmt.Println("Try to create an existing file with flag os.O_EXCL ...")
	_, err = os.OpenFile(filePath1, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	fmt.Printf("error: %v\n", err)
}

type flagDesc struct {
	flag int
	desc string
}
type argDesc struct {
	action string
	flag   int
	perm   os.FileMode
}

//极客时间89练习
func jk89() {
	// 示例1。
	fmt.Printf("The mode for dir:\n%32b\n", os.ModeDir)
	fmt.Printf("The mode for named pipe:\n%32b\n", os.ModeNamedPipe)
	fmt.Printf("The mode for all of the irregular files:\n%32b\n", os.ModeType)
	fmt.Printf("The mode for permissions:\n%32b\n", os.ModePerm)
	fmt.Println()

	// 示例2。
	fileName1 := "something3.txt"
	filePath1 := filepath.Join(os.TempDir(), fileName1)
	fmt.Printf("The file path: %s\n", filePath1)

	argDescList := []argDesc{
		{
			"Create",
			os.O_RDWR | os.O_CREATE,
			0644,
		},
		{
			"Reuse",
			os.O_RDWR | os.O_TRUNC,
			0666,
		},
		{
			"Open",
			os.O_RDWR | os.O_APPEND,
			0777,
		},
	}

	defer os.Remove(filePath1)
	for _, v := range argDescList {
		fmt.Printf("%s the file with perm %o ...\n", v.action, v.perm)
		file1, err := os.OpenFile(filePath1, v.flag, v.perm)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		info1, err := file1.Stat()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("The file permissions: %o\n", info1.Mode().Perm())
	}
}
