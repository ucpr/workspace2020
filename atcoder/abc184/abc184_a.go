package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d\n%d %d", &a, &b, &c, &d)

	fmt.Println(a*d - b*c)
}
