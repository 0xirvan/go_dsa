package main

import "fmt"

// Queue adalah struktur data antrian yang mengikuti prinsip
// FIFO (First in, First Out)
// Elemen pertama yang masuk akan menjadi elemen pertama yang keluar
// https://terapan-ti.vokasi.unesa.ac.id/post/struktur-data-pengertian-fungsi-dan-penerapannya

// Contoh implementasi Queue menggunakan linked list

type Node struct {
	data int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
}

// Time complexity : O(1)
func (q *Queue) Enqueue(data int) {
	newNode := &Node{data: data}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

// Time complexity: O(1)
func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	val := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return val, true
}

func (q *Queue) Display() {
	curr := q.head
	for curr != nil {
		fmt.Printf("[%d]", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	q := &Queue{}
	fmt.Println("Enqueue...")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Display()
	if v, ok := q.Dequeue(); ok {
		fmt.Printf("Deque [%d]\n", v)
	}
}
