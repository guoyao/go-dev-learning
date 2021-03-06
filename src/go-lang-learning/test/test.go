// test
package test

import (
	"fmt"
	"unsafe"
)

type User struct {
	name  string
	email string
	age   uint8
}

var user *User

func (self *User) setName(name string) *User {
	fmt.Println(user == self)
	self.name = name
	fmt.Println(user == self)
	return self
}

func (self User) getName() string {
	return self.name
}

func init() {
	f1 := 1.5
	fmt.Println(&f1)
	fmt.Println(*&f1)

	unsafePtr := unsafe.Pointer(&f1)
	fmt.Println(unsafePtr)
	//fmt.Println(*unsafePtr)

	intPtr := (*int)(unsafePtr)
	fmt.Println(intPtr)
	fmt.Println(*intPtr)

	user = &User{name: "guoyao", email: "wuguoyao@baidu.com"}
	fmt.Println(user.getName())
	user.setName("baby")
	fmt.Println(user.getName())

	arr := [...]int{1, 2, 5}
	slice1 := arr[:2:2]
	fmt.Println(len(arr), cap(arr))
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	slice1 = slice1[:cap(slice1)]
	fmt.Println(len(slice1), cap(slice1), slice1)
	slice1 = arr[:3]
	fmt.Println(slice1)
}
