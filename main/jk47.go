package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//极客时间47章练习,讲述的是 基于HTTP协议的网络服务
// http.Client类型中的Transport代表什么. (单次http调用的过程)
//  IdleConnTimeout: 含义是空闲的链接在多久之后就应该被关闭.
//  DefaultTransport: 会把该字段的值设定为90秒.如果该值为0,那么就标示不关闭空闲链接,可能造成资源泄露.
//  ResponseHeaderTimeout: 从客户端把请求完全递给操作系统到从操作系统那里接收到响应报文头的最大时长.
//  ExpectContinueTimeout: 含义是在客户端递交了请求报文头以后,等待接受第一个响应报文头的时间.
//  TLSHandshakeTimeout: 基于TLS协议的链接在被建立时的握手阶段的超时时间.默认是10.
//  与IdleConnTimeout有关的三个字段.
//   MaxIdleConns  :最大空闲总数.
//   MaxIdleConnsPerHost: 该transport值访问的每一个网络服务的最大空闲连接数,对于某一个transport值访问的每一个
//网络服务,它的空闲连接数都最多只有两个.
//  MaxConnsPerHost: 针对某一个transport值访问的每一个网络服务的最大连接数,不论是否空闲.

// listenAndServer:
//两个问题
// net.Listen函数做了哪些事情?  解析地址端口,然后开始监听
// http.server方法是怎样接受和处理HTTP请求的
// 在一个for循环中,网络监听器的accept方法会不断地调用, 方法返回两个值一个是net.Conn类型,
// 包含了新到来的http请求的网络链接,第二个结果值代表了可能发生的error类型
// accept以后没有报错就会提交到一个新的goroutine里进行执行回调
func main() {

	//simpleHttp()
	simpleHttpServer()

}

func simpleHttp() {
	url1 := "http://baidu.com"
	fmt.Printf("Send request to %q with method GET ...\n", url1)
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Printf("request sending error: %v\n", err)
	}
	defer resp1.Body.Close()
	line1 := resp1.Proto + " " + resp1.Status
	fmt.Printf("The first line of response:\n%s\n", line1)
}

// 监听一个基于TCP协议的网络地址,并对收到的HTTP请求进行处理,
func simpleHttpServer() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 示例1。
	go startServer1(&wg)

	// 示例2。
	go startServer2(&wg)

	wg.Wait()

}
func startServer1(wg *sync.WaitGroup) {
	defer wg.Done()
	var httpServer1 http.Server
	httpServer1.Addr = "127.0.0.1:8080"
	// 由于我们没有定制handler，所以这个网络服务对任何请求都只会响应404。
	if err := httpServer1.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 1 closed.")
		} else {
			log.Printf("HTTP server 1 error: %v\n", err)
		}
	}
}

func startServer2(wg *sync.WaitGroup) {
	defer wg.Done()
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/hi", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/hi" {
			http.NotFound(w, req)
			return
		}
		name := req.FormValue("name")
		if name == "" {
			fmt.Fprint(w, "Welcome!")
		} else {
			fmt.Fprintf(w, "Welcome, %s!", name)
		}
	})
	httpServer2 := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux1,
	}
	if err := httpServer2.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 2 closed.")
		} else {
			log.Printf("HTTP server 2 error: %v\n", err)
		}
	}
}
