package main

import "fmt"

func main() {
	var n, x int
	var s string

	fmt.Scanf("%d %d\n%s", &n, &x, &s)

	for _, c := range s {
		if c == 'o' {
			x += 1
		}
		if c == 'x' && x > 0 {
			x -= 1
		}
	}

	fmt.Println(x)
}
