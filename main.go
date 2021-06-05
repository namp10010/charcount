package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	root := os.Args[1]
	counter := make(map[string]int)
	count := 0
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(),".go"){
			return err
		}

		fmt.Println(path)
		text, err := os.ReadFile(path)

		for _, c := range text {
			counter[string(c)]++
			count++
		}

		return nil
	})

	if _,ok := counter[" "]; ok {
		counter["space"] = counter[" "]
		delete(counter, " ")
	}

	if _,ok := counter["	"]; ok {
		counter["tab"] = counter["	"]
		delete(counter, "	")
	}
	nl := fmt.Sprintf("\n")
	if _,ok := counter[nl]; ok {
		counter["nl"] = counter[nl]
		delete(counter, nl)
	}

	charFreqs := make([]charFreq, 0, len(counter))
	for ch, c := range counter {
		charFreqs = append(charFreqs, charFreq{
			char:  ch,
			count: c,
			freq:  100*float32(c)/float32(count),
		})
	}

	sort.Slice(charFreqs, func(i, j int) bool {
		return charFreqs[i].count > charFreqs[j].count
	})
	for _, cf := range charFreqs {
		fmt.Printf("%v: %.2f%%\n", cf.char, cf.freq)
	}
}

type charFreq struct {
	char string
	count int
	freq  float32
}