package main

import (
	"fmt"
	"errors"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("panic error")

}


func test1(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {z = 0}
		}()
		z = x / y
		return
	}()
	fmt.Println("x / y =", z)
}

type error interface {
	Error() string
}

var errDivByZero = errors.New("division by 0")

func div(x, y int) (int, error) {
	if y == 0 {return 0, errDivByZero}
	return x / y, nil
}


func main() {
	test()
	test1(1, 0)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("got an error:", err)
		}
	}()

	switch z, err := div(10, 0); err {
	case nil:
		fmt.Println(z)
	case errDivByZero:
		panic(err)
	}
	fmt.Println("go on")
}

