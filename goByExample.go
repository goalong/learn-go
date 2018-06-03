package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")

	// 字符串可以用+拼接
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("86 / 4.0 = ", 86 / 4.0)

	// 布尔操作
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// 关键字var声明一个或多个变量，没有声明初始值的将被当作该类型的0值取值, := 也可以声明并初始化一个变量
	var a = "haha"
	var b bool
	var c, d int
	fmt.Println(a, b, c, d)

	// 常量
	const ok = 100
	fmt.Println(ok)

	// for循环,有几种形式
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// if else
	num := 14
	if num > 10 {
		fmt.Println("num bigger than 10")
	} else {
		fmt.Println("num smaller than 10")
	}

	// switch
	i = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// array, 拥有固定长度且类型相同的序列
	var l [5]int
	fmt.Println(l)

	// slice, 切片， 长度可变的数组
	s := make([]int, 4)
	s[2] = 3
	fmt.Println(s)
	s = s[2:]
	fmt.Println(s)

	// map, 无序的键值对
	m := make(map[int]string)
	m[2] = "真二"
	fmt.Println(m)

	// ranage，使用range来遍历
	for i, v := range s {
		fmt.Println(i, v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	r := plus(2, 9)
	fmt.Println(r)

	_s := sum(1,2,3)

	nums := []int {1,2,3,4}
	_s = sum(nums...)
	fmt.Println(_s)

	// 指针
	t := 10
	t_p := &t
	fmt.Println(t, t_p)
	*t_p = 15
	fmt.Println(t, t_p)

	// struct,
	p := person{name: "john", age: 20}
	fmt.Println(p.name, p.age)

	// method, 用于结构的方法
	_r := rect{width:9, height:6}
	_rp := &_r
	fmt.Println(_rp.area())

	// interface, 接口， 包含一系列的方法, 要实现一个接口需要实现这个接口的所有方法
	_c := circle{radius: 5}
	measure(_c)

}

func plus(a int, b int) int {
	return a+b
}

// 接收多个参数
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


