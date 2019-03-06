package main

import (
	"awesomeProject/golang-base/common"
	"awesomeProject/golang-base/common/op"
	"errors"
	"fmt"
	"os"
	"runtime/pprof"
)

//极客时间48章,关于程序性能分析有关的基础知识
// runtime中有一些底层的API,可以被用来收集或者输出Go程序运行过程中的一些关键指标.
// 生成的关键指标,需要标准工具的支持 go tool pprof  和 go tool trace
// 分析程序性能的概要文件有三种: CPU概要文件,内存概要文件,阻塞概要文件. 这些文件都是pprof序列化后的二进制文件. 我需要工具才可以查看

func main() {
	simpleCPUProfile()
}

var (
	profileName = "cpuprofile.out"
)

//CPU概要信息进行采样
//go tool pprof cpuprofile.out
func simpleCPUProfile() {
	//这里需要指定文件的生成路径,其实我完全可以用接受参数的方法来进行,或者rest
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	if err := startCPUProfile(f); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}
	//模拟高负载的搞错
	if err = common.Execute(op.CPUProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	stopCPUProfile()
}
func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
