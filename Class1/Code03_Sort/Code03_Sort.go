package main

func selectionSort(arr []int) {
	// Your code goes here
	if arr == nil || len(arr) < 2 {
		return
	}

	// 0~N-1
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		// i~N-1上找最小值下标
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

}

func main() {

}
