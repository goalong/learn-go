package main

import "fmt"

// 接口习惯以er结尾，只有方法签名，没有实现，没有数据字段，允许嵌套，类型可以实现多个接口
// 任何类型的方法集中只要包含一个接口的全部方法，就表示它实现了该接口

type Stringer interface {
	String() string
}


type Printer interface {
	Stringer  // 接口的嵌入
	Print()
}


type User struct {
	id int
	name string
}


func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}


func (self *User) Print() {
	fmt.Println(self.String())
}


func main() {
	var t Printer = &User{1, "jack"}
	t.Print()
}