package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

//极客时间26章 讲sync.Mutex和 sync.RWMutex,虽然golang是用通讯的方式来共享数据,但是它依然提供了易用的同步工具

//笔记1: 使用互斥锁时有哪些注意事项
// 1. 不要重复锁定互斥所
// 2. 不要忘记解锁互斥锁,必要时使用defer语句
// 3. 不要对尚未锁定或者已解锁的互斥锁解锁
// 4. 不要在多个函数之间直接传递互斥锁 这里是 值传递,没有用的小老弟
var me sync.Mutex   //实现的locker接口
var rw sync.RWMutex //实现的locker接口
// 笔记2: 读写锁的规则
//  1.在写锁已经被锁定的情况下再试图锁定写锁,会阻塞当前的goroutine
//  2.在写锁已经被锁定的情况下试图锁定读锁,会阻塞当前的goroutine
//  3.在读锁已经被锁定的情况下试图锁定写锁,会阻塞当前的goroutine
//  4.在读锁已经被锁定的情况下在试图锁定读锁,并不会 阻塞当前的goroutine

//另一个方面,对写锁解锁成功,会唤醒所有因试图锁定读锁而被阻塞的goroutine.并且,这通常会使它们都成功完成对读锁的锁定
// 还有,对读锁进行解锁,只会在没有其他读锁锁定的前提下,唤醒 因试图锁定写锁,而被阻塞的goroutine,并且
// 最终只会有一个被唤醒的goroutine能够成功的对写锁锁定,其他的goroutine还要原地等待.
func main() {
	//这部分涉及的知识类似与java中的lock锁
	//metux2()
	//练习3
	metux3()
}

//摘选自极客时间的demo58
// 普通锁配合多goroutine练习
//  锁 其实在goroutine的共享数据区中才有存在的意义与价值
func metux1() {
	var protecting uint = 1
	// buffer 代表缓冲区。
	var buffer bytes.Buffer

	const (
		max1 = 5  // 代表启用的goroutine的数量。
		max2 = 10 // 代表每个goroutine需要写入的数据块的数量。
		max3 = 10 // 代表每个数据块中需要有多少个重复的数字。
	)

	// mu 代表以下流程要使用的互斥锁。
	var mu sync.Mutex
	// sign 代表信号的通道。
	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				//程序读空结构体然后阻塞的defer实现
				sign <- struct{}{}
			}()
			for j := 1; j <= max2; j++ {
				// 准备数据。
				header := fmt.Sprintf("\n[id: %d, iteration: %d]",
					id, j)
				data := fmt.Sprintf(" %d", id*j)
				// 写入数据。
				if protecting > 0 {
					mu.Lock()
				}
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}

	for i := 0; i < max1; i++ {
		<-sign
	}
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}

// singleHandler 代表单次处理函数的类型。
type singleHandler func() (data string, n int, err error)

// handlerConfig 代表处理流程配置的类型。
type handlerConfig struct {
	//得益于 golang中可以把函数当成属性传递
	handler   singleHandler // 单次处理函数。
	goNum     int           // 需要启用的goroutine的数量。
	number    int           // 单个goroutine中的处理次数。
	interval  time.Duration // 单个goroutine中的处理间隔时间。
	counter   int           // 数据量计数器，以字节为单位。
	counterMu sync.Mutex    // 数据量计数器专用的互斥锁。

}

// count 会增加计数器的值，并会返回增加后的计数。
func (hc *handlerConfig) count(increment int) int {
	hc.counterMu.Lock()
	defer hc.counterMu.Unlock()
	hc.counter += increment
	return hc.counter
}
func metux2() {
	// mu 代表以下流程要使用的互斥锁。
	// 在下面的函数中直接使用即可，不要传递。
	var mu sync.Mutex

	// genWriter 代表的是用于生成写入函数的函数。
	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			// 准备数据。
			data = fmt.Sprintf("%s\t",
				time.Now().Format(time.StampNano))
			// 写入数据。
			mu.Lock()
			defer mu.Unlock()
			fmt.Println("写完了")
			n, err = writer.Write([]byte(data))
			return
		}
	}

	// genReader 代表的是用于生成读取函数的函数。
	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}
			// 读取数据。
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	// buffer 代表缓冲区。
	var buffer bytes.Buffer

	// 数据写入配置。
	writingConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}
	// 数据读取配置。
	readingConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    10,
		number:   2,
		interval: time.Millisecond * 100,
	}

	// sign 代表信号的通道。
	sign := make(chan struct{}, writingConfig.goNum+readingConfig.goNum)

	// 启用多个goroutine对缓冲区进行多次数据写入。
	for i := 1; i <= writingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= writingConfig.number; j++ {
				time.Sleep(writingConfig.interval)
				data, n, err := writingConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := writingConfig.count(n)
				log.Printf("writer [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// 启用多个goroutine对缓冲区进行多次数据读取。
	for i := 1; i <= readingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= readingConfig.number; j++ {
				time.Sleep(readingConfig.interval)
				var data string
				var n int
				var err error
				for {
					data, n, err = readingConfig.handler()
					if err == nil || err != io.EOF {
						break
					}
					// 如果读比写快（读时会发生EOF错误），那就等一会儿再读。
					time.Sleep(readingConfig.interval)
				}
				if err != nil {
					log.Printf("reader [%d-%d]: error: %s",
						i, j, err)
					continue
				}
				total := readingConfig.count(n)
				log.Printf("reader [%d-%d]: %s (total: %d)",
					i, j, data, total)
			}
		}(i)
	}

	// signNumber 代表需要接收的信号的数量。
	signNumber := writingConfig.goNum + readingConfig.goNum
	// 等待上面启用的所有goroutine的运行全部结束。
	for j := 0; j < signNumber; j++ {
		<-sign
	}

}

//示例二

// counter 代表计数器。
type counter struct {
	num uint         // 计数。
	mu  sync.RWMutex // 读写锁。
}

// number 会返回当前的计数。
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

// add 会增加计数器的值，并会返回增加后的计数。
func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}
func count(c *counter) {
	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	go func() { // 用于增加计数。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 1, j)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 2, k)
		}
	}()
	<-sign
	<-sign
	<-sign
}

//我加的读锁 一定需要解读锁,不能来解写锁
//这是个读写锁示例,我先开的goroutine 在里面休眠1秒然后解读锁,主goroutine会加读锁加写锁 这样不会永远阻塞的!
// 如果我的goroutine里面写的是解写锁 那就错了 因为 加读锁以后再加写锁就GG了永远阻塞了
func redundantUnlock() {
	var rwMu sync.RWMutex
	//var r sync.Mutex

	// 示例1。
	//rwMu.Unlock() // 这里会引发panic。

	// 示例2。
	//rwMu.RUnlock() // 这里会引发panic。

	//加读锁
	go func() {
		fmt.Println("解读锁")
		time.Sleep(time.Second)
		rwMu.RUnlock()
	}()
	rwMu.RLock()
	rwMu.Lock()
	time.Sleep(time.Second * 2)
	//解写锁
	//rwMu.Unlock()
	//rwMu.RUnlock()
	//rwMu.Unlock() // 这里会引发panic。
	//rwMu.RUnlock()

	// 示例4。
	//rwMu.Lock()
	//rwMu.RUnlock() // 这里会引发panic。
	//rwMu.Unlock()
}
func metux3() {
	//c := counter{}
	//count(&c)
	redundantUnlock()
}
