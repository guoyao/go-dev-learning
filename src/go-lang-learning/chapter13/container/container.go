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
	"hash/crc32"
	"io/ioutil"
	"crypto/sha1"
)

func Main() {
//	link()
//	sortTest()
//	crc()
//	isTowFileSame()
	shaTest()
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

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortTest() {
	kids := []Person{
		{"Jill",9},
		{"Jack",10},
	}
	fmt.Println(kids)
	sort.Sort(ByName(kids))
//	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

func crc() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func isTowFileSame() {
	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test1.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

func shaTest() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
