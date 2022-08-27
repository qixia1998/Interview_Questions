package main

import "fmt"

// 堆排序
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}

func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i // 假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rigthchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { // 防止越过边界
			topmax = leftchild
		}
		if rigthchild <= length-1 && arr[rigthchild] > arr[topmax] { // 防止越过边界
			topmax = rigthchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}
