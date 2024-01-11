package ListTraverse

import "fmt"

type Node struct {
	data int
	next *Node
}

func createLL(arr []int) *Node {
	first := Node{
		arr[0],
		nil,
	}
	last := &first
	for _, v := range arr[1:] {
		t := &Node{v, nil}
		last.next = t
		last = t
	}
	return &first
}

func ListTraverse(reversed bool) {
	arr := []int{3, 5, 7, 10, 15}
	node := createLL(arr)
	if reversed {
		rtraverse(node)
	} else {
		traverse(node)
	}
}

func traverse(p *Node) {
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}

func rtraverse(p *Node) {
	if p != nil {
		rtraverse(p.next)
		fmt.Println(p.data)
	}
}
