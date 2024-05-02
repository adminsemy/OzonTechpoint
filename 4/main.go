package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Data struct {
	id   uint32
	time uint32
}

func scan() {
	// file, _ := os.Open("./236/50")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t uint16
	fmt.Fscan(input, &t)
	var n uint32
	var athlets []Data
	var index uint32
	var result []uint32
	for i := uint16(0); i < t; i++ {
		fmt.Fscan(input, &n)
		athlets = make([]Data, n)
		result = make([]uint32, n)
		for j := uint32(0); j < n; j++ {
			athlets[j].id = j
			fmt.Fscan(input, &athlets[j].time)
		}
		sort.Slice(athlets, func(i, j int) bool {
			return athlets[i].time < athlets[j].time
		})
		index = 1
		result[athlets[0].id] = index
		for j := uint32(1); j < n; j++ {
			index++
			if athlets[j].time-athlets[j-1].time > 1 {
				result[athlets[j].id] = index
			} else {
				result[athlets[j].id] = result[athlets[j-1].id]
			}
		}
		for j := uint32(0); j < n; j++ {
			fmt.Fprint(out, result[j], " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	scan()
}
