package main

import "fmt"

// Linked list merupakan struktur data yang terdiri dari serangkaian elemen yang terhubung satu
// sama lain menggunakan reference / pointer
// Setiap elemen dalam linked list disebut dengan node
// Node pertama dalam linked list disebut dengan head
// Node terakhir dalam linked list disebut dengan tail
// https://terapan-ti.vokasi.unesa.ac.id/post/memahami-konsep-dan-jenis-jenis-linked-list-dalam-struktur-data

type Node struct {
	// datanya, kita buat int saja biar simpel
	data int
	// pointer ke node selanjutnya
	next *Node
}

type SinglyLinkedList struct {
	// dalam singly linked list biasanya kita hanya perlu menyimpan head nya saja
	// meskipun tailnya juga bisa di simpan
	head *Node
}

// Time complexity: O(1)
func (l *SinglyLinkedList) PushFront(data int) {
	newNode := &Node{data: data}
	newNode.next = l.head
	l.head = newNode
}

// Time complexity: O(n) karena harus cari tail dulu
func (l *SinglyLinkedList) PushBack(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

// Time complexity: O(n) Linear search node
func (l *SinglyLinkedList) Find(target int) *Node {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.data == target {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

// TIme complexity: O(1) Tapi perlu cari node dulu
// Jadi nya O(1) + proses cari node O(n) = O(n) jadinya
func (l *SinglyLinkedList) InsertAfter(node *Node, data int) {
	if node == nil {
		return
	}
	newNode := &Node{data: data}
	newNode.next = node.next
	node.next = newNode
}

// Time complexity: O(1)
func (l *SinglyLinkedList) DeleteFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

// Time complexity: O(1)
// Tapi hari cari dulu node nya, sama kek fungsi insertafter
func (l *SinglyLinkedList) DeleteAfter(node *Node) {
	if node == nil || node.next == nil {
		return
	}
	node.next = node.next.next
}

func (l *SinglyLinkedList) Display() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func main() {
	list := &SinglyLinkedList{}
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

	node = list.Find(3)
	list.DeleteAfter(node)
	fmt.Printf("setelah menghapus node setelah node %d: ", node.data)
	list.Display()
}
