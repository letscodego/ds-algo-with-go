package array

import "errors"

type CustomArray struct {
	size int
	item []int
}

func (a *CustomArray) NewCustomArray(size int) {
	a.item = make([]int, size)
	a.size = 0
}

func (a *CustomArray) Insert(item int) {
	if len(a.item) > a.size {
		a.item[a.size] = item
		a.size++
	} else {
		temp := a.item
		a.item = make([]int, len(a.item)*2)
		copy(a.item, temp)
		a.item[a.size] = item
		a.size++
	}
}

func (a *CustomArray) RemoveAt(index int) error {
	if index > a.size || index < 0 {
		return errors.New("invalid index")
	}
	a.item = append(a.item[:index], a.item[index+1:]...)
	a.size--
	return nil
}

func (a CustomArray) IndexOf(item int) int {

	index := indexOf(a.item, item)
	return index
}

func indexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}
