package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if strings.HasSuffix(d.Name(), ext) {
			a = append(a, s)
		}
		return nil
	})
	return a
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
	err = scanner.Err()
	return
}

func main() {
	fields := make(map[string]int)
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
	for field, count := range fields {
		fmt.Println(count, field)
	}
}
