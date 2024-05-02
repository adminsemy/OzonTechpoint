package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Data struct {
	Id    uint32
	Value uint32
}

func scan() {
	// file, _ := os.Open("./233/3")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m uint32
	fmt.Fscan(input, &n, &m)
	friends := make([]Data, n)
	result := make([]uint32, n)
	for i := uint32(0); i < n; i++ {
		friends[i].Id = i
		fmt.Fscan(input, &friends[i].Value)
	}
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].Value > friends[j].Value
	})
	for i := uint32(0); i < n; i++ {
		if friends[i].Value >= m {
			fmt.Fprintln(out, -1)
			return
		}
		result[friends[i].Id] = m
		m--
	}
	for i := uint32(0); i < n; i++ {
		fmt.Fprint(out, result[i], " ")
	}
	fmt.Fprintln(out)
}

func main() {
	scan()
}
