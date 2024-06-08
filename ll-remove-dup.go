package main

import "fmt"

// remove duplicates from unsorted linked list
// How would you solve this problem if temp buffer is not allowed?

//[] 				-> []
//[5] 			    -> [5]
//[5, 2, 5, 8, 5]   -> [5, 2, 8]
//[9, 9, 9] 		-> [9]
//[1, 2, 5, 2, 1]   -> [1, 2, 5]

type Node struct {
	next *Node
	data int
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

func (ll *Node) Append(data int) *Node {
	itr := ll
	for itr.next != nil {
		itr = itr.next
	}
	itr.next = NewNode(data)
	return ll
}

func (ll *Node) RemoveDuplicatesWithBuff() {
	uniques := map[int]int{}
	itr := ll
	prev := itr
	for itr != nil {
		if _, ok := uniques[itr.data]; ok {
			prev.next = itr.next
		} else {
			uniques[itr.data] = 1
			prev = itr
		}
		itr = itr.next
	}
}

func (ll *Node) RemoveDuplicatesWithSelection() {
	selected := ll
	for selected != nil {
		itr, prev := selected, selected
		for itr != nil {
			if itr.data == selected.data {
				prev.next = itr.next
			} else {
				prev = itr
			}
			itr = itr.next
		}
		selected = selected.next
	}
}

func (ll *Node) Display() {
	itr := ll
	for itr != nil {
		fmt.Print(itr.data, " ")
		itr = itr.next
	}
	fmt.Println()
}

func main() {
	//[5] 			    -> [5]
	ll := NewNode(5)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()

	//[5, 2, 4] 			    -> [5, 2, 4]
	ll = NewNode(5).Append(2).Append(4)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()

	//[9, 9, 9] 		-> [9]
	ll = NewNode(9).Append(9).Append(9)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()

	//[1, 2, 5, 2, 1]   -> [1, 2, 5]
	ll = NewNode(1).Append(2).Append(5).Append(2).Append(1)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()

	//[5, 2, 5, 8, 5]   -> [5, 2, 8]
	ll = NewNode(5).Append(2).Append(5).Append(8).Append(5)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()

	//[5, 2, 2, 1, 3, 4]   -> [5, 2, 1, 3, 4]
	ll = NewNode(5).Append(2).Append(2).Append(1).Append(3).Append(4)
	//ll.RemoveDuplicatesWithBuff()
	ll.RemoveDuplicatesWithSelection()
	ll.Display()
}
