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

type freqs []freq

func (f freqs) Len() int { return len(f) }

func (f freqs) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func (f freqs) Less(i, j int) bool {
	if f[i].Count != f[j].Count {
		// Sort by count descending
		return f[i].Count > f[j].Count
	}

	// Then by key ascending
	return f[i].Key < f[j].Key
}

func main() {
	fields := make(map[string]int)
	fieldsf := freqs{}

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

	sort.Sort(fieldsf)

	for _, f := range fieldsf {
		fmt.Printf("%2d %s\n", f.Count, f.Key)
	}
}
