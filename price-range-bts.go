package main

import (
	"fmt"
)

type tree struct {
	left  *tree
	value int
	right *tree
}

func new(val int) *tree {
	t := &tree{nil, val, nil}
	return t
}

func (t *tree) insert(val int) {
	insertRec(t, val)
}

func insertRec(tr *tree, val int) *tree {
	if tr == nil {
		return &tree{nil, val, nil}
	}
	if val < tr.value {
		tr.left = insertRec(tr.left, val)
	}
	if val > tr.value {
		tr.right = insertRec(tr.right, val)
	}
	return tr
}

func productsInRange(t *tree, c chan int, low int, high int) {
	walker(t, c, low, high)
	close(c)
}

func walker(t *tree, c chan int, low int, high int) {
	if t == nil {
		return
	}

	if low <= t.value && t.value <= high {
		c <- t.value
	}

	if low <= t.value {
		walker(t.left, c, low, high)
	}

	if t.value <= high {
		walker(t.right, c, low, high)
	}
}

func main() {
	bst := new(9)
	bst.insert(6)
	bst.insert(14)
	bst.insert(20)
	bst.insert(1)
	bst.insert(30)
	bst.insert(8)
	bst.insert(17)
	bst.insert(5)

	ch := make(chan int)

	low, high := 7, 20

	go productsInRange(bst, ch, low, high)

	fmt.Printf("products in the price range between %d and %d:\n", low, high)
	for i := range ch {
		println(i)
	}

}
