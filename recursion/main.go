package main

import "fmt"

// Recursion it's just a function yang memanggil diri nya sendiri
// menyelesaikan masalah yang lebih kecil dari masalah awal,
// sampai kondisi tertentu.

// Struktur rekursif
// Base case : kondisi berhenti
// Recursif case: pemanggilan diri sendiri dengan input yang makin dekat dengan base case

// Call Stack / Stack Frame
// Setiap pemanggilan fungsi disimpan di stack memory (nilai param, variabel lokal, posisi return)

// Contoh penggunaan rekursif
// Faktorial: perkalian bilangan tersebut dengan semua bilangan bulat positif di bawahnya hingga 1
// n! = n × (n-1) × (n-2) × ... × 1
// 0! = 1

func Faktorial(n int) int {
	if n < 0 {
		panic("faktorial tidak didefinisikan untuk bilangan negatif")
	}

	// Base case
	if n == 0 || n == 1 {
		return 1
	}

	return n * Faktorial(n-1)
}

// Fibonnaci seq: barisan yang setiap sukunya merupakan penjumlahan dari dua suku sebelumnya.
// Formula: F(n) = F(n-1) + F(n-2)

func Fibonnaci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonnaci(n-1) + Fibonnaci(n-2)
}

func main() {
	fmt.Println("5! = ", Faktorial(5))
	fmt.Println("Fib(5) = ", Fibonnaci(6))
}
