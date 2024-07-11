package benchmark

// 不预分配内存
func MakeSliceWithoutAlloc() []int {
	var newSlice []int
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}

// MakeSliceWithPreAlloc 通过预分配的Slice的存储空间构造

func MakeSliceWithPreAlloc() []int {
	var newSlice []int
	newSlice = make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}
	return newSlice
}
