// M.DAVI ILYAS RENALDO
package main

import "fmt"

func selectionSort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func menghitungMedian(data []int) {
	n := len(data)
	var median int
	if n%2 == 0 {
		median = (data[n/2-1] + data[n/2]) / 2
	} else {
		median = data[n/2]
	}
	fmt.Println(median)
}

func main() {
	var arrData []int
	var input int

	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			selectionSort(arrData)
			menghitungMedian(arrData)
		} else {
			arrData = append(arrData, input)
		}
	}
}