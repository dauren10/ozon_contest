package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n) // read the number of pairs

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)  // read two integers
		fmt.Fprintln(out, a+b) // print the sum
	}
}
