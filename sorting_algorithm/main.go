package main

import "fmt"

// Insertion sort adalah algoritma sorting yang membangun array terurut
// secara bertahap dengan cara mengambil satu elemen dari bagian yang
// belum terurut dan menyisipkan nya ke posisi yang tepat pada bagian
// yang sudah terurut.
// awalnya belum semua data terurut
// setiap langkah menambahkan 1 elemen ke bagian yang terurut
// Contoh:
// Langkah 0: [5] | [3,4,1]
// Langkah 1: [3,5] | [4,1]
// Langkah 2: [3,4,5] | [1]
// Langkah 3: [1,3,4,5]

// Time complexity:
// Best case: O(n) Jika array sudah ter sort
// Average/Worst case: O(n^2)
func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		temp := s[i]
		j := i - 1

		for j >= 0 && s[j] > temp {
			s[j+1] = s[j]
			j--
		}

		s[j+1] = temp
	}
}

func main() {
	arr := []int{2, 8, 1, 12, 3, 11, 10, 6, 4, 9, 5, 7}
	InsertionSort(arr)
	fmt.Println(arr)
}
