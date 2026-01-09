package main

// https://neetcode.io/problems/dynamicArray

type DynamicArray struct {
	data     []int
	size     int
	capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.data[i] = n
}
func (da *DynamicArray) Pushback(n int) {
	if da.capacity == da.size {
		da.Resize()
	}

	da.data[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	da.size--
	return da.data[da.size]
}

func (da *DynamicArray) Resize() {
	newCap := da.capacity * 2
	newData := make([]int, newCap)

	for i := 0; i < da.size; i++ {
		newData[i] = da.data[i]
	}
	da.capacity = newCap
	da.data = newData
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
