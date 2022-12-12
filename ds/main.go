package main

import (
	"ds-algo/ds/array"
	"ds-algo/ds/linkedList"
	"fmt"
)

func main() {

	var m array.CustomArray
	m.NewCustomArray(3)
	m.Insert(10)
	m.Insert(20)
	m.Insert(30)
	m.Insert(40)
	fmt.Println(m)
	m.RemoveAt(-1)
	fmt.Println(m)
	fmt.Println(m.IndexOf(30))

	var l linkedList.LinkedList
	l.AddFirst(1)
	l.AddLast(10)
	l.AddLast(20)
	l.AddLast(50)
	l.AddLast(60)
	fmt.Println(l.Size)
	fmt.Println(l.ToArray())

	fmt.Println(l.GetKthFromTheEnd(2))

}
