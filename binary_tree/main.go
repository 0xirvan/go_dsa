package main

import "fmt"

// Binary tree adalah struktur data non-linear berbentuk pohon.
// Setiap node memiliki maksimal dua anak
// Left Child (kiri)
// Right Child (kanan)

// Komponen binary tree
// Node: Terdiri dari data/value pointer ke left and right child
// Root: Node paling atas (Tidak memiliki parent)
// Leaf: Node tanpa anak (Pointer kiri kanan nya nil)
// Subtree: Pohon kecil yang berasal dari suatu node

// Properti binary tree
// Height: tinggi pohon, di hitung dari root sampai leaf terdalam (jumlah level − 1)
// Depth: jarak suatu node dari root (milik node, bukan tree)
// Degree: jumlah anak dari suatu node

// Contoh:
//        A
//       / \
//      B   C
//     / \
//    D   E

// Traversal
// Cara mengunjungi semua node pada binary tree satu per satu dengan urutan tertentu.

// Preorder: Root -> kiri -> kanan
// Baca A -> Ke Kiri -> Baca B -> Ke kiri -> Baca D -> D Gak punya child -> Balik ke B -> Ke kanan B
// -> Baca E -> E selesai -> Balik Ke B -> Balik Ke A -> Ke kanan A -> Baca C -> Selesai

// Inorder: Kiri -> Root -> Kanan (Jangan baca node saat pertama datang, baca setelah kiri selesai)
// Mulai dari A -> Ke kiri -> B -> Ke kiri -> D -> Kiri D Kosong -> Baca D -> Balik ke B
// -> Baca B -> Ke kanan -> E -> Kiri E kosong -> Baca E -> Balik ke B -> Balik ke A
// -> Baca A -> Ke Kanan -> C -> Kiri C kosong -> Baca C

// Postorder: Kiri -> Kanan -> Baca (Node di baca paling terakhir, baca saat mau di tinggal)
// Mulai dari A -> Ke Kiri -> B -> Ke kiri -> D -> Kiri D kosong -> kanan D Kosong -> Baca D
// Balik Ke B -> Ke Kanan -> E -> Kiri E kosong -> Kanan E kosong -> Baca E -> Balik Ke B
// kiri dan kanan B sudah selesai -> Baca B -> Balik Ke A -> Ke kanan -> C -> Kiri C kosong
// -> Kanan C kosong -> Baca C -> Balik ke  A -> Kiri kanan A sudah Selesai -> Baca A

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.value, " ")
	PreorderTraversal(node.left)
	PreorderTraversal(node.right)
}

func InorderTraversal(node *Node) {
	if node == nil {
		return
	}
	InorderTraversal(node.left)
	fmt.Print(node.value, " ")
	InorderTraversal(node.right)
}

func PostorderTraversal(node *Node) {
	if node == nil {
		return
	}
	PostorderTraversal(node.left)
	PostorderTraversal(node.right)
	fmt.Print(node.value, " ")
}

func main() {
	root := NewNode(1)           // A
	root.left = NewNode(2)       // B
	root.right = NewNode(3)      // C
	root.left.left = NewNode(4)  // D
	root.left.right = NewNode(5) // E

	fmt.Print("Preorder Travesal: ")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Print("Inorder Travesal: ")
	InorderTraversal(root)
	fmt.Println()

	fmt.Print("Postorder Travesal: ")
	PostorderTraversal(root)
	fmt.Println()
}
