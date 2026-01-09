package main

// https://neetcode.io/problems/singlyLinkedList

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	idxCounter := 0
	currentNode := ll.head

	for currentNode != nil {
		if idxCounter == index {
			return currentNode.data
		}
		idxCounter++
		currentNode = currentNode.next
	}

	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	newNode := &Node{data: val}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) InsertTail(val int) {
	newNode := &Node{data: val}

	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList) Remove(index int) bool {
	if ll.head == nil {
		return false
	}

	if index == 0 {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		return true
	}

	currentNode := ll.head
	for i := 0; i < index-1 && currentNode.next != nil; i++ {
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		return false
	}

	if currentNode.next == ll.tail {
		ll.tail = currentNode
	}

	currentNode.next = currentNode.next.next
	return true
}

func (ll *LinkedList) GetValues() []int {
	var res []int

	currentNode := ll.head
	for currentNode != nil {
		res = append(res, currentNode.data)
		currentNode = currentNode.next
	}
	return res
}
