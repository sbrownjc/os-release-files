package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func find(root, ext string) (result []string) {
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if strings.HasSuffix(d.Name(), ext) {
			result = append(result, s)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking dir: %s", err)
	}

	return result
}

func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		err = fmt.Errorf("error scanning: %w", scanner.Err())
	}

	return lines, err
}

type freq struct {
	Key   string
	Count int
}

func main() {
	fields := make(map[string]int)
	fieldsf := []freq{}

	for _, s := range find(".", "os-release") {
		lines, err := FileToLines(s)
		if err != nil {
			fmt.Println(err)
		}

		for _, l := range lines {
			split := strings.SplitN(l, "=", 2)
			if split[0] != "" {
				fields[split[0]]++
			}
		}
	}

	for f, c := range fields {
		fieldsf = append(fieldsf, freq{Key: f, Count: c})
	}

	sort.Slice(fieldsf, func(i, j int) bool {
		if fieldsf[i].Count != fieldsf[j].Count {
			// Sort by count descending
			return fieldsf[i].Count > fieldsf[j].Count
		}

		// Then by key ascending
		return fieldsf[i].Key < fieldsf[j].Key
	})

	for _, f := range fieldsf {
		fmt.Printf("%2d %s\n", f.Count, f.Key)
	}
}
