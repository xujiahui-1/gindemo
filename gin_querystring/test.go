package main

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	dummy := &Node{
		Val:  -1,
		Next: nil,
		Prev: nil,
	}
	dummy.Next = dummy
	dummy.Prev = dummy
	return MyLinkedList{
		dummy: dummy,
	}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index != 0 {
		return -1
	}
	return head.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy.Next,
		Prev: dummy,
	}
	dummy.Next.Prev = node
	dummy.Next = node

}

func (this *MyLinkedList) AddAtTail(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Next: dummy,
		Prev: dummy.Prev,
	}
	dummy.Prev.Next = node
	dummy.Prev = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index > 0 {
		return
	}
	node := &Node{
		Val:  val,
		Next: head,
		Prev: head.Prev,
	}
	head.Prev.Next = node
	head.Prev = node
}
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.dummy.Next == this.dummy {
		return
	}
	head := this.dummy.Next
	//找到要删除的节点head
	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}
	if index == 0 {
		head.Next.Prev = head.Prev
		head.Prev.Next = head.Next
	}
}
