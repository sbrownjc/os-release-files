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

// Find searches a given directory root for files matching the given extension,
// and returns a list of filenames.
func Find(root, ext string) (result []string) {
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.Type().IsRegular() && strings.HasSuffix(d.Name(), ext) {
			result = append(result, s)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error walking dir: %s", err)
	}

	return result
}

// FileToLines takes a given file and splits it into lines.
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

// WholePercentage determines the percentage of the num over the denom,
// and returns an integer, unrounded.
func WholePercentage(num, denom int) int {
	const oneHundred = 100

	return int((float64(num) / float64(denom) * oneHundred))
}

// Frequency is a struct for counting the number of instances of a given Key
type Frequency struct {
	Key   string
	Count int
}

func main() {
	files := 0
	fields := make(map[string]int)
	fieldsf := []Frequency{}

	for _, s := range Find(".", "os-release") {
		files++

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
		fieldsf = append(fieldsf, Frequency{Key: f, Count: c})
	}

	sort.Slice(fieldsf, func(i, j int) bool {
		if fieldsf[i].Count != fieldsf[j].Count {
			return fieldsf[i].Count > fieldsf[j].Count
		}

		return fieldsf[i].Key > fieldsf[j].Key
	})

	for _, f := range fieldsf {
		fmt.Printf("%2d %s (%d%%)\n", f.Count, f.Key, WholePercentage(f.Count, files))
	}
}
