package main

import "fmt"

func fibonacci() func() int {
	f0, f1 := 0, 1

	return func() (x int) {
		x = f0
		f0, f1 = f1, f0+f1
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
