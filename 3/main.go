package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan() {
	// file, _ := os.Open("./242/1")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, q uint32
	fmt.Fscan(input, &n, &q)
	users := make([]uint32, n+1)
	var t, id uint32
	var index uint32
	var result uint32
	for i := uint32(0); i < q; i++ {
		fmt.Fscan(input, &t, &id)
		if t == 1 {
			index++
			users[id] = index
		}
		if t == 2 {
			result = users[id]
			if result < users[0] {
				result = users[0]
			}
			fmt.Fprintln(out, result)
		}
	}
}

func main() {
	scan()
}
