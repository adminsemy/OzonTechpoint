package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []*Data  `json:"folders"`
}

func check(data *Data, isVirus bool) uint32 {
	result := uint32(0)
	for i := 0; i < len(data.Files); i++ {
		if len(data.Files[i]) > 4 && data.Files[i][len(data.Files[i])-5:] == ".hack" {
			isVirus = true
		}
		result++
	}
	if !isVirus {
		result = 0
	}
	for i := 0; i < len(data.Folders); i++ {
		result += check(data.Folders[i], isVirus)
	}
	return result
}

func scan() {
	// file, _ := os.Open("./72/24")
	// defer file.Close()
	// input := bufio.NewReader(file)
	input := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t uint16
	fmt.Fscan(input, &t)
	var n uint16
	var data []string
	var str string
	var result uint32
	for i := uint16(0); i < t; i++ {
		fmt.Fscan(input, &n)
		data = make([]string, n+1)
		for j := uint16(0); j <= n; j++ {
			str, _ = input.ReadString('\n')
			data[j] = str
		}
		str := strings.Join(data, "")
		files := new(Data)
		json.Unmarshal([]byte(str), &files)
		result = check(files, false)
		fmt.Fprintln(out, result)
	}
}

func main() {
	scan()
}
