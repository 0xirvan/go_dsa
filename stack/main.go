package main

import (
	"fmt"
)

// Stack adalah struktur data tumpukan yang mengikuti prinsip
// LIFO (Last in, First out)
// Elemen yang terakhir ditambahkan akan menjadi yang pertama di hapus
// https://terapan-ti.vokasi.unesa.ac.id/post/struktur-data-pengertian-fungsi-dan-penerapannya

// Contoh implementasi stack menggunakan linked list

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
}

// Time complexity: O(1)
func (s *Stack) Push(data int) {
	newNode := &Node{data: data}
	newNode.next = s.head
	s.head = newNode
}

// Time complexity: O(1)
func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	val := s.head.data
	s.head = s.head.next
	return val, true
}

// Time complexity: O(1)
func (s *Stack) Peek() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	return s.head.data, true
}

// Time complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("Stack kosong")
		return
	}
	curr := s.head
	for curr != nil {
		fmt.Printf("[%d]\n", curr.data)
		curr = curr.next
	}
}

func main() {
	s := &Stack{}

	fmt.Println("Menambahkan 1 ke stack: ")
	s.Push(1)
	s.Display()

	fmt.Println("Menambahkan 2 ke stack: ")
	s.Push(2)
	s.Display()

	fmt.Println("Menambahkan 3 ke stack: ")
	s.Push(3)
	s.Display()

	fmt.Println("Menghapus elemen dari stack:")
	s.Pop()
	s.Display()

	fmt.Println("Menghapus elemen dari stack:")
	s.Pop()
	s.Display()

	fmt.Println("Melihat elemen: ")
	if v, ok := s.Peek(); ok {
		fmt.Printf("[%d]\n", v)
	}
}
