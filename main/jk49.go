package main

import (
	"awesomeProject/golang-base/common"
	"awesomeProject/golang-base/common/op"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
)

var (
	profileName1   = "memprofile.out"
	memProfileRate = 8
)

var (
	profileName2     = "blockprofile.out"
	blockProfileRate = 2
	debug            = 0
)

//极客时间49章
func main() {

	//simpleMemProfile()
	//simplieBlock()
	log.Println(http.ListenAndServe("localhost:8082", nil))
}

func simpleMemProfile() {

	f, err := common.CreateFile("", profileName1)
	if err != nil {
		fmt.Printf("memory profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	startMemProfile()
	if err = common.Execute(op.MemProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	// 收集好的内存概要信息,写到指定的写入器中.
	if err := stopMemProfile(f); err != nil {
		fmt.Printf("memory profile stop error: %v\n", err)
		return
	}
}

func startMemProfile() {
	//设定内存概要信息采样频率
	// 平均每分配8个字节,就对堆内存的使用情况进行一次采样 ,缺省是512kb
	runtime.MemProfileRate = memProfileRate
}

func stopMemProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.WriteHeapProfile(f)
}

// 获取到阻塞的概要信息
func simplieBlock() {

	f, err := common.CreateFile("", profileName2)
	if err != nil {
		fmt.Printf("block profile creation error: %v\n", err)
		return
	}
	defer f.Close()
	startBlockProfile()
	if err = common.Execute(op.BlockProfile, 10); err != nil {
		fmt.Printf("execute error: %v\n", err)
		return
	}
	if err := stopBlockProfile(f); err != nil {
		fmt.Printf("block profile stop error: %v\n", err)
		return
	}
}
func startBlockProfile() {
	//只要发生阻塞事件的持续时间达到了2纳秒.就进行采样
	runtime.SetBlockProfileRate(blockProfileRate)
}

func stopBlockProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	// debug 0 内存地址  1包名 函数名 源码路径 代码行号
	// name : goroutine  heap  allocs  threadcreate  block mutex
	return pprof.Lookup("block").WriteTo(f, debug)
}

//http://localhost:8082/d/pprof/symbol
func HttpRe() {
	mux := http.NewServeMux()
	pathPrefix := "/d/pprof/"
	mux.HandleFunc(pathPrefix,
		func(w http.ResponseWriter, r *http.Request) {
			name := strings.TrimPrefix(r.URL.Path, pathPrefix)
			if name != "" {
				pprof.Handler(name).ServeHTTP(w, r)
				return
			}
			pprof.Index(w, r)
		})
	mux.HandleFunc(pathPrefix+"cmdline", pprof.Cmdline)
	mux.HandleFunc(pathPrefix+"profile", pprof.Profile)
	mux.HandleFunc(pathPrefix+"symbol", pprof.Symbol)
	mux.HandleFunc(pathPrefix+"trace", pprof.Trace)

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}
	server.ListenAndServe()
}
