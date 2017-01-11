package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
	"stringutil"
	"sync"
	"sync/atomic"
	"time"
)

var mutex = &sync.Mutex{} //可简写成：var mutex sync.Mutex
var total_tickets int32 = 10

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH") + "\n")
	fmt.Printf("浮点数：%f\n", math.Pi)
	fmt.Printf("十六进制：%X\n", 10)
	var x int = 100
	fmt.Printf("整型：%d\n", x)

	var name2 string = "你好"
	fmt.Printf("name: %s\n", name2)

	y := 100
	fmt.Printf("y:%d\n", y)

	const pi float32 = 3.1415923
	fmt.Printf("pi:%f\n", pi)
	fmt.Println("sfasdf")

	var array [5]int
	array[2] = 23
	fmt.Println("array length:", len(array))
	fmt.Println("array:", array)

	a := [5]int{0, 1, 2, 3, 4}
	b := a[1:]
	fmt.Println(b)

	if y%2 == 0 {
		fmt.Printf("100被2整除了")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("for循环:", i)
	}

	c := 5
	pointD := &c
	fmt.Println("指针练习1：", pointD, *pointD)

	*pointD = 2
	fmt.Println("指针练习2：", pointD, *pointD)

	c = 3
	fmt.Println("指针练习3：", pointD, *pointD)

	var d *[]int = new([]int)
	var e []int = make([]int, 10)
	fmt.Println("指针练习4：", d, e)
	*d = make([]int, 10, 10)
	fmt.Println("指针练习5：", (*d)[2])

	fmt.Println("函数SUM：", sum(5, 11))

	f := make(map[string]int)
	f["one"] = 1
	fmt.Println("字典：", f)

	g, success := multi_ret("three")
	if success {
		fmt.Println("字典取值成功：", g, success)
	} else {
		fmt.Println("字典取值失败：", g, success)
	}

	person := Person{"Tom", 30, "vincentchow0213@gmail.com"}
	fmt.Println("结构体：", person)

	h := rect{3.0, 4.0}
	i := circle{2}
	j := []shape{&h, &i}
	fmt.Println("接口使用：", j[0].perimeter())

	k, l := CopyFile("E:\\Vincent\\work_go\\source.txt", "E:\\Vincent\\work_go\\target.txt")
	fmt.Println("IO处理：", k, l)

	for i := 0; i < 5; i++ {
		// defer fmt.Printf("DEFER 回收资源：%d ", i)
	}

	//生成随机种子
	// rand.Seed(time.Now().Local().Unix())
	//并发处理
	runtime.GOMAXPROCS(4)
	// var name string
	// for i := 0; i < 3; i++ {
	// 	name = fmt.Sprintf("go_%02d", i) //生成ID
	// 	//生成随机等待时间，从0-4秒
	// 	go routine(name, time.Duration(rand.Intn(5))*time.Second)
	// }
	//让主进程停住，不然主进程退了，goroutine也就退了
	// var input string
	// fmt.Println("输入任意值返回结果")
	// fmt.Scanln(&input)

	//进销存管理，需要加锁
	// for i := 0; i < 5; i++ {
	// 	go sell_tickets(i)
	// }
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println(total_tickets, "done")

	//原子操作
	var cnt uint64 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				// atomic.AddUint32(&cnt, 1)
				atomic.AddUint64(&cnt, 1)
			}
		}()
	}
	time.Sleep(time.Second)             //等一秒钟等goroutine完成
	cntFinal := atomic.LoadUint64(&cnt) //取数据
	fmt.Println("atom count:", cntFinal)

	//信道Channel(1是buff)
	channel := make(chan string, 1)
	channel2 := make(chan string)

	go func() {
		channel <- "hello"
		time.Sleep(3 * time.Second)
		channel <- "World"
		channel <- "!"
	}()
	go func() {
		channel2 <- "坏人"
	}()
	select {
	case msg1 := <-channel:
		fmt.Println("channel1 received", msg1)
	case msg2 := <-channel2:
		fmt.Println("channel2 received", msg2)
	}
	// msg1 := <-channel
	// fmt.Println(msg1)
	// msg2 := <-channel
	// fmt.Println(msg2)
	// msg3 := <-channel
	// fmt.Println(msg1, msg2, msg3)

	var input string
	fmt.Scanln(&input)
}

func sell_tickets(i int) {
	for total_tickets > 0 {
		mutex.Lock()
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets--
			fmt.Println(i, total_tickets)
		}
		mutex.Unlock()
	}
}

func routine(name string, delay time.Duration) {

	t0 := time.Now()
	fmt.Println(name, " start at ", t0)

	time.Sleep(delay)

	t1 := time.Now()
	fmt.Println(name, " end at ", t1)

	fmt.Println(name, " lasted ", t1.Sub(t0))
}

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func sum(a int, b int) int {
	return a + b
}

func multi_ret(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2}
	var success bool
	var value int

	value, success = m[key]
	return value, success
}

type Person struct {
	name  string
	age   int
	email string
}

//---------- 接 口 --------//
type shape interface {
	area() float64      //计算面积
	perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type rect struct {
	width, height float64
}

func (r *rect) area() float64 { //面积
	return r.width * r.height
}

func (r *rect) perimeter() float64 { //周长
	return 2 * (r.width + r.height)
}

//----------- 圆  形 ----------//
type circle struct {
	radius float64
}

func (c *circle) area() float64 { //面积
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { //周长
	return 2 * math.Pi * c.radius
}
