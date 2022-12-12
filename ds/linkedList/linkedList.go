package linkedList

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	First *Node
	Last  *Node
	Size  int
}

func (n *Node) NewNode(value int, nextNode *Node) {
	n.Value = value
	n.Next = nextNode
}

func (l *LinkedList) AddFirst(value int) {
	var f Node
	f.NewNode(value, nil)

	if l.Size == 0 {
		l.First = &f
		l.Last = &f
	} else {
		f.Next = l.First
		l.First = &f
	}
	l.Size++
}

func (l *LinkedList) AddLast(value int) {
	var last Node
	last.NewNode(value, nil)

	if l.Size == 0 {
		l.First = &last
		l.Last = &last
	} else {
		l.Last.Next = &last
		l.Last = &last
	}
	l.Size++
}

func (l *LinkedList) DeleteFirst() {
	if l.Size <= 1 {
		l.First = nil
		l.Last = nil
		l.Size = 0
		return
	}
	temp := l.First.Next
	l.First.Next = nil
	l.First = temp
	l.Size--
}

func (l *LinkedList) DeleteLast() {
	if l.Size <= 1 {
		l.First = nil
		l.Last = nil
		l.Size = 0
		return
	}

	node := l.First.Next
	for node != nil {
		if node.Next == l.Last {
			node.Next = nil
			l.Last = node
		}
		node = node.Next
	}
	l.Size--
}

func (l *LinkedList) Contains(value int) bool {
	return l.IndexOf(value) != -1
}

func (l *LinkedList) IndexOf(value int) int {
	node := l.First
	index := 0
	for node != nil {
		if node.Value == value {
			return index
		}
		node = node.Next
		index++
	}
	return -1
}

func (l LinkedList) ToArray() []int {
	item := make([]int, l.Size)
	if l.Size == 0 {
		return item
	}
	current := l.First
	index := 0
	for current != nil {
		item[index] = current.Value
		current = current.Next
		index++
	}
	return item
}

func (l LinkedList) ReverseWithNewList() LinkedList {
	if l.Size <= 1 {
		return l
	}

	var reverse LinkedList
	var current = l.First
	for current.Next != nil {
		reverse.AddFirst(current.Value)
		current = current.Next
	}
	reverse.AddFirst(l.Last.Value)
	return reverse
}

func (l *LinkedList) Reverse() {
	if l.Size <= 1 {
		return
	}

	previous := l.First
	current := l.First.Next

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	l.Last = l.First
	l.Last.Next = nil
	l.First = previous
}

func (l LinkedList) GetKthFromTheEnd(k int) (int, error) {
	if k <= 0 {
		return -1, fmt.Errorf("input value can't be less than 1")
	}

	a := l.First
	b := l.First
	for i := 0; i < k-1; i++ {
		b = b.Next
		if b == nil {
			return -1, fmt.Errorf("input value is greater than list size")
		}
	}

	for b != l.Last {
		a = a.Next
		b = b.Next

	}
	return a.Value, nil
}
