package main

import (
	"fmt"
)

func main() {
	var x int
	var list []int
	fmt.Println("Enter a Number")
	fmt.Scanf("%d", &x)
	fmt.Println(x)
	for i:=1; i<x; i++ {
		if checkPrime(i) {
			list = append(list, i)
		}
	}
	fmt.Println(list)
}

func checkPrime(num int) bool{
	if num == 1 || num == 0 {
		return false
	} else {
		for i :=2 ; i< num/2; i++ {
			if num%i == 0 {
				return false
			}
		}		
	}
	return true
}
