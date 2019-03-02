package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync/atomic"
)

//极客时间30讲,关于golang语言的atomic的下半场!

//CAS乐观锁的一种,配合for循环实现自旋锁,不用我多BB了吧
//衍生问题4: 假设我已经对保证了对一个变量的写操作都是原子操作,比如:加或减,存储,交换等那么我的对它进行读取的时候,还有必要使用原子操作吗?
// 回答: 此处而已对比读写锁,读锁和写锁是互斥锁,对一个共享资源的保护,就需要在写和读的层面做到完全的保护

//笔记
// 原子值对Store方法对其参数值(被存储值)有两个强制的约束,一个约束是,参数值不能为nil(接口例外,接口的组成比较特出,值类型也属于接口接口,除非接口没有任何实现)
//另一个约束是,参数值的类型不能与首个被存储的值的类型不同.
// 建议: 不要对外暴露原子变量,不要传递原子值及其指针值,也不要用引用类型
func main() {

	//第二个
	//hq()
	//avalueT()
	//slice()
	//atomicF()
	structAtomic()
}

type atom struct {
	Name string
}

//原子值测试
func avalueT() {
	var at atomic.Value
	at2 := &atom{"qws"}
	at23 := &atom{"qws23"}
	// 原子值存储的结构一定要前后一致,否则会panic
	at.Store(at2)
	at.Store(at23)
	//向后加载的值为准
	fmt.Println(at.Load())

}

//原子值赋值等
func hq() {
	var at atomic.Value
	at23 := atom{"qws23"}
	// 原子值存储的结构一定要前后一致,否则会panic
	// 我将值传递给atomic.value存储,这里是值复制.,那么我从atomic.Value取值的时候就不是at23而是其副本,我对他的修改也是对副本的修改
	at.Store(at23)
	a24 := at23
	atome := at.Load().(atom)
	atome.Name = "111"
	a24.Name = "222"
	//24值传递 跟你们没屌关系
	fmt.Println(a24)
	//我还是原来的23
	fmt.Println(at23)
	//你已经不是原来的atome
	fmt.Println(atome)

}

// 原子值切片的操作
func slice() {
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box6.Store(replica)
	}
	store(v6)
	// v6的本质是我复制了一个切片,这里对v6切片对我store的切片没有任何的影响,
	v6[2] = 5
	//无影响
	fmt.Println(box6.Load())
	//打印v6的值
	fmt.Println(v6)

}

//摘选自git-极客时间24课时代码

func atomicF() {
	// 示例4。
	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("Store an error with message %q to box4.\n", v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Store a value of the same type to box4.")
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Printf("Store a value of type %T that implements error interface to box4.\n", v42)
		// 虽然都实现了error接口,但是他们的结构体类型不一样
		//box4.Store(v42) // 这里会引发一个panic，报告存储值的类型不一致。
	}
	fmt.Println()
}
func ato2() {
	var box atomic.Value
	v1 := []int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Println("Copy box to box3.")
	//复制以后就是一个新值了
	box3 := box // 原子值在真正使用后不应该被复制！
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	ints := box.Load().([]int)
	ints[0] = 0
	// 复制的是同一个底层的切片  有点恶心啊
	fmt.Println(box.Load())
	fmt.Println(box3.Load())
}

//这是我自己的结构体,蛮有意思的
func structAtomic() {
	v4 := errors.New("something wrong")
	v41 := io.EOF
	// v42跟 v4和v41是不同的!
	v42 := interface{}(&os.PathError{}).(error)
	// 示例5。
	box5, err := NewAtomicValue(v4)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same type to box5.")
	err = box5.Store(v41)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("Store a value of type %T that implements error interface to box5.\n", v42)
	err = box5.Store(v42)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println()

}

//推荐用法,里面多一层验证还是比较安全的
type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

//这是我自己的store方法,注意,里面做了安全校验,感觉还蛮不错的
func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}
