package ListTraverse

import "fmt"

func ListLinearSearch() {
	arr := []int{3, 5, 7, 10, 25, 8, 32, 2}
	head := createLL(arr)
	temp, head := search(head, head, 8)
	if temp != nil {
		fmt.Println("Key is found:", temp.data)
	} else {
		fmt.Println("Key not found")
	}
	traverse(head)
}

func search(first, p *Node, key int) (*Node, *Node) {
	var q *Node
	for p != nil {
		if key == p.data {
			q.next = p.next
			p.next = first
			first = p
			return p, first
		}
		q = p
		p = p.next
	}
	return nil, first
}
