package main

import (
	"fmt"
	"os"
	"strconv"
)

func CountLenght(n int) int {
	if n<0 {
		return 0
	}
	if n==0 {
		return 1
	}
	lenght := 0
	for ; n != 0; n = n / 10 {
		lenght++
	}
	return lenght
}

func SumNum(n int) int {
	if n<0 {
		return 0
	}
	summ := 0
	count := CountLenght(n)
	for i := 0; i < count; i++ {
		summ += n % 10
		n = n / 10
	}
	return summ
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	//fmt.Scanf("%d", &n)
	fmt.Println(CountLenght(n))
	fmt.Println(SumNum(n))
}
