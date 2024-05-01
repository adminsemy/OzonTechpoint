package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan() {
	// file, _ := os.Open("./238/33")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var s string
	fmt.Fscan(input, &s)
	result := []byte(s)
	var n uint16
	fmt.Fscan(input, &n)
	var start, end uint16
	var r string
	var index uint16
	for i := uint16(0); i < n; i++ {
		fmt.Fscan(input, &start, &end, &r)
		index = 0
		for j := start - 1; j < end; j++ {
			result[j] = r[index]
			index++
		}
	}
	fmt.Fprintln(out, string(result))
}

func main() {
	scan()
}
