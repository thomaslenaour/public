package main

import (
	solution "../../solutions"
	"github.com/01-edu/z01"
)

type Node2 = NodeL
type List2 = solution.List
type NodeS2 = solution.NodeL
type ListS2 = List

func listPushBackTest2(l *ListS2, l1 *List2, data interface{}) {
	n := &Node2{Data: data}
	n1 := &NodeS2{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		iterator := l.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}

	if l1.Head == nil {
		l1.Head = n1
	} else {
		iterator1 := l1.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n1
	}
}

func main() {
	link := &List2{}
	link2 := &ListS2{}
	table := []solution.NodeTest{}
	table = solution.ElementsToTest(table)
	table = append(table,
		solution.NodeTest{
			Data: []interface{}{"Hello", "man", "how are you"},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBackTest2(link2, link, arg.Data[i])
		}
		aux := solution.ListSize(link)
		aux2 := ListSize(link2)
		if aux != aux2 {
			z01.Fatalf("ListSize(%v) == %d instead of %d\n", solution.ListToString(link.Head), aux2, aux)
		}
		link = &List2{}
		link2 = &ListS2{}
	}
}