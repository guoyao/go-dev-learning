/**
 * User: guoyao
 * Time: 06-27-2013 17:47
 * Blog: http://www.guoyao.me
 */

package container

import (
	"fmt"
	"container/list"
	"sort"
)

func Main() {
//	link()
	sortTest()
}

func link() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}
}


type Person struct {
	Name string
	Age int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortTest() {
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
	}
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
