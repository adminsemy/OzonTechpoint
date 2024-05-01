package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan() {
	// file, _ := os.Open("./235/1")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int16
	fmt.Fscan(input, &t)
	var a, b int16
	for i := int16(0); i < t; i++ {
		fmt.Fscan(input, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}

func main() {
	scan()
}
