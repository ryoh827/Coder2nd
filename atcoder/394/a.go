package main

import (
	"fmt"
	"strings"
)

// https://atcoder.jp/contests/abc394/tasks/abc394_a
func main() {
	var S string
	fmt.Scan(&S)

	c := strings.Count(S, "2")
	fmt.Println(strings.Repeat("2", c))
}
