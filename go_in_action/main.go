package main

import ("fmt"
		"runtime"
		"sync"
		"sync/atomic")

var (
	counter int64
	wg sync.WaitGroup
)

func main() {
	s1 := []int {2,1,3,5,6}
	s2 := s1[2:4]
	s1[2] = 10
	// 两个切片共享同一个底层数组，一个切片修改共享部分，另一个切片也会被修改
	fmt.Println(s1, s2)

	s3 := make([]int, 3, 3)
	s3[0] = 2
	s3[2] = 5

	// 在函数间传递一个map并不会对map进行一份拷贝，函数中对这个map的修改会在函数外也生效，和切片类似
	// 将切片或者map传递给函数成本很小
	m := make(map[int]int)
	m[12] = 12
	modifyMap(m)
	fmt.Println(m)  // map多了（100， 100）这个键值对


	// 数值，字符串，布尔，对这些类型的值进行增加或删除时，会创建一个新值，传递给函数或方法时
	// 传递的是对应值的副本

	// Go语言的引用类型有切片，映射，通道，接口，和函数类型，声明以上类型的变量时，创建的变量被称作标头值(header),
	// 标头值里包含一个指向底层数据结构的指针，因此通过复制传递一个引用类型的值的副本，本质上是在共享底层数据结构


	// 在接口上调用方法时，接收者是值的方法可以通过指针调用，因为指针会被先解引用
	// 接受者是指针的方法不可以通过值调用，因为接口中的值没有地址

	//user1 := user{name:"jack", email:"jack@ab.com"}

	// 小写字母开头的变量是私有的，意味着在其他包中不可见，大写的是公开的，可以在其他包中使用

	// 映射的键可以是任何值，切片，函数及包含切片的结构由于具有引用语义不能作为键

	dict := map[string]string{}
	dict["a"] = "ac"
	value, exists := dict["a"]  // 从映射获取值并判断键是否存在
	if exists {fmt.Println(value)}
	delete(dict, "a")
	fmt.Println(dict)
	// 函数间传递一个映射时并不会对该映射进行一份拷贝，函数对这个映射修改时，所有对这个映射的引用都会察觉到修改

	// 接口是声明了一组行为并支持多态的类型

	// 嵌入类型提供了扩展类型的能力，而无需使用继承

	// 原子函数
	wg.Add(2) // 表示等待两个goroutine
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("final counter:", counter)



	

}


func modifyMap(data map[int]int) {
	data[100] = 100
}

type user struct {
	name string
	email string
}

type admin struct {
	user
	level string
}


func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}