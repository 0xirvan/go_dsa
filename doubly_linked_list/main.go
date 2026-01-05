package main

import "fmt"

// Berbeda dengan singly linked list, doubly linked list menyimpan
// reference / pointer dari node sebelumnya, jadi kek dua arah
type Node struct {
	data int
	// Node sebelumnya
	prev *Node
	// Node selanjutnya
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	// Ada tail agar operasi insert dari belakang nya mempunyai time complexity O(1)
	tail *Node
}

// Time complexity: O(1)
func (l *DoublyLinkedList) PushFront(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	oldHead := l.head
	newNode.next = oldHead
	oldHead.prev = newNode
	l.head = newNode
}

// Time complexity: O(1)
func (l *DoublyLinkedList) PushBack(data int) {
	newNode := &Node{data: data}

	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		return
	}

	oldTail := l.tail
	newNode.prev = oldTail
	oldTail.next = newNode
	l.tail = newNode
}

// Time complexity: O(n) Linear search
func (l *DoublyLinkedList) Find(target int) *Node {
	currentNode := l.head

	for currentNode != nil {
		if currentNode.data == target {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

// Time complexity: O(n), tapi harus cari node target dulu
func (l *DoublyLinkedList) InsertAfter(node *Node, data int) {
	if node == nil {
		return
	}

	newNode := &Node{data: data}
	newNode.prev = node
	newNode.next = node.next

	if node.next != nil {
		node.next.prev = newNode
	} else {
		l.tail = newNode
	}

	node.next = newNode
}

// Time complexity: O(n)
func (l *DoublyLinkedList) DeleteFront() {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.next
	l.head.prev = nil
}

// Time complexity: O(n)
func (l *DoublyLinkedList) DeleteBack() {
	if l.tail == nil {
		return
	}

	if l.tail == l.head {
		l.tail = nil
		l.head = nil
		return
	}

	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *DoublyLinkedList) Display() {
	currentNode := l.head
	fmt.Print("nil -> ")
	for currentNode != nil {
		fmt.Printf("%d <-> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func main() {
	list := &DoublyLinkedList{}
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)
	list.PushFront(5)

	fmt.Print("Menampilkan list awal: ")
	list.Display()

	list.PushBack(1)
	fmt.Print("Setelah menambah node di akhir: ")
	list.Display()

	node := list.Find(3)
	list.InsertAfter(node, 10)
	fmt.Printf("Setelah menambah node setelah %d: ", node.data)
	list.Display()

	list.DeleteFront()
	fmt.Print("Setelah menghapus front: ")
	list.Display()

	list.DeleteBack()
	fmt.Print("Setelah menghapus back: ")
	list.Display()
}
