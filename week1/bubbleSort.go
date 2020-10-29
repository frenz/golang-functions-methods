package main

import (
	"fmt"
	"os"
)

func printError(err error){
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func Swaps(slice []int, index int ){
	if index +1 < len(slice) {
		aux := slice[index]
		slice[index] = slice[index+1]
		slice[index+1] = aux
	}
}

func Compare(slice []int, first, second int) int{
	if slice[first] <= slice[second]{
		return -1
	}
	return 1
}

func BubbleSort(slice []int){
	for  i:= 0; i < Size; i++ {
		for j := 0; j < Size -i - 1; j++ {
			if Compare(slice, j , j+1) == 1 {
				Swaps(slice, j)
			} 

		}
	}
}
const Size = 10

func main() {
	slice := make([]int, Size)
	fmt.Println("Go course - bubbleSort.go ...")
	for i := 0; i < Size; i++ {
		var input int
		fmt.Printf("Insert int %d of %d \n", i+1, Size)
		_, err := fmt.Scanf("%d", &input)
		printError(err)
		slice[i] = input
	}
	BubbleSort(slice)
	fmt.Println(slice)
	os.Exit(0)
}
