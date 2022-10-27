package main

import (
	"fmt"
)

func main() {
	a := []int{}
	printSlice(a)
	a = append(a, 1)
	printSlice(a)
	a = append(a, 2)
	printSlice(a)
	a = append(a, 3)
	printSlice(a)
	a = append(a, 4)
	printSlice(a)
	a = append(a, 5)
	printSlice(a)
}

type FilterFunc func(string) bool

func printSlice(arr []int) {
	fmt.Printf("%p / length: %v / capacity: %v / %v \n", arr, len(arr), cap(arr), arr)
}

func meinFilter(s []string, filter FilterFunc) []string {
	out := []string{}
	for _, str := range s {
		if filter(str) {
			out = append(out, str)
		}
	}
	return out
}

func laengenFilter(laenge int) FilterFunc {
	return func(s string) bool {
		return len(s) > laenge
	}
}
