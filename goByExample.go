package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"math/rand"
	"sort"
	"strings"
	"encoding/json"
	"os"
	"strconv"
	"net/url"
	"net"
	"crypto/sha1"
	"encoding/base64"
	"os/signal"
	"syscall"
	"reflect"
)

// go by example 的一些代码例子，并做了注释便于理解

const a = "I am a const value"
var pr = fmt.Println
func main() {
	p := 2000  // 变量的声明和赋值， var 或者 :=
	fmt.Println("hello world" + "2018")  //  字符串可以用 + 拼接
	fmt.Println(a)
	fmt.Println(p)

	// for循环，有几种形式,
	// for init; condition; post { }
	// for condition { }
	for i:= 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// if else
	if num:=9;num<0{
		fmt.Println(num,"is negative")
	}else if num<10{
		fmt.Println(num,"has 1 digit")
	}else{
		fmt.Println(num,"has multiple digits")
	}

	// switch 语句
	v := 3
	switch  {
	case v % 2 == 0:
		fmt.Println("v is even")
	case v % 2 == 1:
		fmt.Println("v is odd")
	}

	// 数组, array, 拥有固定长度且类型相同的序列
	a := [5]int {1,2,3,4}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	// slice, 切片， 长度可变的数组
	b := a[2:]
	fmt.Println(b)

	// map, 键值对, 无序
	c := make(map[int]string)
	c[6] = "he"
	c[8] = "oy"
	fmt.Println(c)

	// range, 来遍历可迭代的对象，如数组，切片，map等
	for k, v := range c {
		fmt.Println(k, v)
	}

	// 函数，使用普通变量作为函数参数的时候，在传递参数时只是对变量值得拷贝，
	// 即将实参的值复制给变参，当函数对变参进行处理时，并不会影响原来实参的值
	i := 8
	zeroval(i)
	fmt.Println(i)

	// 函数的变量不仅可以使用普通变量，还可以使用指针变量，
	// 使用指针变量作为函数的参数时，在进行参数传递时将是一个地址，即将实参的内存地址复制给变参，这时对变参的修改也将会影响到实参的值
	zeroptr(&i)
	fmt.Println(i)

	// 结构
	d := rect {2, 5}
	//  area是rect结构的方法
	rs := d.area()
	fmt.Println(rs)

	/*var t inf1 = rect {4,6}
	fmt.Println(t.area())

	f("direct")
	//要想让这个函数在goroutine中触发，使用go f(s)。这个新的goroutine将会与调用它的并行执行
	go f("goroutine")
	//我们也可以启动一个调用匿名函数的goroutine
	go func(msg string){
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done") */

	// 通道, 用于连接不同的goroutine
	messages := make(chan string)
	// 通过channel <- syntax的语法来为channel传值
	go func() {messages <- "ping"}()

	//  <-channel的语法来获取channel中的值
	msg := <- messages
	fmt.Println(msg)

	// Channel Buffering, 一个通道可以接收多个值
	messages = make(chan string, 3)
	messages <- "haha"
	messages <- "nihao"
	messages <- "世界"
	fmt.Println(<-messages, <-messages, <-messages)

	done := make(chan bool)

	// Channel Synchronization, 可以使用channel来同步goroutine的执行
	go worker(done)
	// 这里会阻塞知道从上面的goroutine里获取到channel的值
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "lavala")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// select and timeout
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// select 可以用于等待多个goroutine操作
	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	// select case 和timeout的应用，看几个情形哪个先完成，执行先完成的
	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 3")
	}

	// non-blocking channel operation

	messagers := make(chan string)
	signals := make(chan bool)

	// 用select case  default 进行非阻塞的接收
	select {
	case msg := <- messagers:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no msg received")
	}

	// 非阻塞的接收，多个通道
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// closing channel
	jobs := make(chan int, 5)
	finish := make(chan bool)

	// 这个worker持续接收jobs传来的值
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job",j)
			} else {
				fmt.Println("received all jobs")
				finish <- true
				return
			}}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close之后上面的worker接受到more为false
	close(jobs)
	fmt.Println("sent all jobs")
	<-finish


	// range来遍历channels

	queue := make(chan string, 2)
	queue <- "1"
	queue <- "2"
	close(queue)

	// 通道关闭之后仍能收到它接收的值
	for elm := range queue {
		fmt.Println(elm)
	}


	// timer，代表未来的一个事件

	timer1 := time.NewTimer(time.Second * 2)

	//
	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	// timer 可以取消
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stoped")
	}

	// ticker，用于每隔一定事件重复执行
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	// ticker可以停止
	ticker.Stop()

	// worker pools，例子这里是三个worker来共同完成五个工作

	jobs = make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动三个goroutine
	for w:= 1; w <= 3; w++ {
		go work(w, jobs, results)
	}

	// 这里为jobs通道发送值
	for j:= 1; j <= 5; j++ {
		jobs <-j
	}
	//关闭jobs通道
	close(jobs)

	// 获取results通道收到的五个值
	for a:= 1; a<=5; a++ {
		<-results
	}

	// rate limiting

	limiter := make(chan time.Time, 3)
	for i:=0; i< 3; i++ {
		limiter <- time.Now()
	}
	go func() {
		for t:=range time.Tick(time.Millisecond*200) {
			limiter <- t
		}
	}()

	requests := make(chan int, 10)
	for i:=1; i<=10; i++ {
		requests <- i
	}
	close(requests)
	// 前三个request会被立即处理，因为limiter里已经有三个值，之后limiter每200毫秒接收一个值，所以request也每200毫秒处理一个
	for req:= range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// atomic counter
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()

	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)


	// stateful goroutines
	var readOps uint64 = 0
	var writeOps uint64 = 0
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r:=0; r< 100; r++ {
		go func() {
			for {
				read := &readOp {
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:= 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	// sorting
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings: ", strs)

	ints := []int {3,5, 1}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)


	// sorting by functions

	fruits := []string {"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// panic
	//panic("some problen")
	//_, err := os.Create("/tmp/file")
	//if err != nil {
	//	panic(err)
	//}


	// collection functions

	strs = []string {"peach", "apple", "banana", "plum"}
	fmt.Println(Index(strs, "apple"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))


	// string functions

	pr("Contains:  ", strings.Contains("test", "es"))
	pr("Count:     ", strings.Count("test", "t"))
	pr("HasPrefix: ", strings.HasPrefix("test", "te"))
	pr("HasSuffix: ", strings.HasSuffix("test", "st"))
	pr("Index:     ", strings.Index("test", "e"))
	pr("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	pr("Repeat:    ", strings.Repeat("a", 5))
	pr("Replace:   ", strings.Replace("foo", "o", "0", -1))
	pr("Replace:   ", strings.Replace("foo", "o", "0", 1))
	pr("Split:     ", strings.Split("a-b-c-d-e", "-"))
	pr("ToLower:   ", strings.ToLower("TEST"))
	pr("ToUpper:   ", strings.ToUpper("test"))
	pr()


	// json

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int {"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page: 1,
		Fruits: []string {"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte (`{"nums": 6.13, "strs": ["a", "b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)


	//_num := dat["num"].(float64)
	//fmt.Println(_num)

	_strs := dat["strs"].([]interface{})

	str1 := _strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	_d := map[string]int {"apple": 5, "lettuce": 7}
	enc.Encode(_d)

	// unix time
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now, secs, nanos)
	fmt.Println(time.Unix(secs, 0))


	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339,"2012-11-01T22:08:41+00:00")
	fmt.Println(t1)

	// number parsing
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	_i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(_i)

	k, _ := strconv.Atoi("125")
	fmt.Println(k)


	// url parsing
	_s := "postgres://user:pass@host.com:4321/index?k=v#f"
	u, err := url.Parse(_s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme, u.User, u.User.Username())
	pw, _ := u.User.Password()
	fmt.Println(pw)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host, port, u.Path, u.Fragment)
	fmt.Println(u.RawQuery)

	// sha1 hash
	_s = "sha1 this string"
	h:= sha1.New()
	h.Write([]byte(_s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// base64 encoding
	_q := "hello rocket"
	eq := base64.StdEncoding.EncodeToString([]byte(_q))
	fmt.Println(eq)

	// signals
	sigs := make(chan os.Signal, 1)
	done = make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")







}

func zeroval(i int) {
	i = 0
}


func zeroptr(i *int) {
	*i = 0
}

type rect struct {
	width, height float64
}

// 方法，
func (r rect) area() float64 {
	return r.width * r.height
}


type inf1 interface {
	area() float64

}


func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,";",i)
	}
}


func worker(done chan bool) {
	fmt.Println("starting...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true

}

// channel用于函数参数时，可以指定是用于接收值还是发送值, chan<- 接收，<-chan 发送
func ping(pings chan<- string, msg string) {
	pings <- msg
}


func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func work(id int, jobs <- chan int, results chan<- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}


type readOp struct {
	key int
	resp chan int
}


type writeOp struct {
	key int
	val int
	resp chan bool
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}


func Index(vs []string, t string) int {
	for i, v:= range vs {
		if v == t{
			return i
		}

	}
	return -1
}


func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}


func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}


func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}


func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}


func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v:= range vs {
		vsm[i] = f(v)
	}
	return vsm
}


type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}




