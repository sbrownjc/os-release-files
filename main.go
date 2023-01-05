package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"golang.org/x/exp/slices"
)

// Keys contained in the official spec: https://www.freedesktop.org/software/systemd/man/os-release.html
var specKeys = []string{
	"ANSI_COLOR",
	"ARCHITECTURE",
	"BUG_REPORT_URL",
	"BUILD_ID",
	"CPE_NAME",
	"DEFAULT_HOSTNAME",
	"DOCUMENTATION_URL",
	"HOME_URL",
	"ID_LIKE",
	"ID",
	"IMAGE_ID",
	"IMAGE_VERSION",
	"LOGO",
	"NAME",
	"PORTABLE_PREFIXES",
	"PRETTY_NAME",
	"PRIVACY_POLICY_URL",
	"SUPPORT_END",
	"SUPPORT_URL",
	"SYSEXT_LEVEL",
	"SYSEXT_SCOPE",
	"VARIANT_ID",
	"VARIANT",
	"VERSION_CODENAME",
	"VERSION_ID",
	"VERSION",
}

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
	Name  string
	Count int
	OSes  map[string]struct{}
}

// FieldsForFile parses an os-release file and returns the values for ID, VERSION_ID,
// and the names of all fields contained within.
func FieldsForFile(path string) (osid string, version string, fields []string) {
	lines, err := FileToLines(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, l := range lines {
		split := strings.SplitN(l, "=", 2)
		key := split[0]
		if key != "" {
			fields = append(fields, key)
			value := strings.Trim(strings.TrimSpace(split[1]), `"`)

			if key == "ID" {
				osid = value
			}
			if key == "VERSION_ID" {
				version = value
			}
		}
	}

	return osid, version, fields
}

// MapKeys takes a set map and returns the keys in a sorted order
func MapKeys(m map[string]struct{}) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	allOses := map[string]struct{}{}
	fileCount := 0
	fields := []string{}
	fieldsfreq := map[string]Frequency{}

	for _, s := range Find(".", "os-release") {
		fileCount++
		var osid string
		osid, _, filefields := FieldsForFile(s)
		for _, f := range filefields {
			var ff Frequency
			var ok bool
			if ff, ok = fieldsfreq[f]; !ok {
				fields = append(fields, f)
				ff = Frequency{Name: f, OSes: map[string]struct{}{}}
			}
			ff.Count++
			ff.OSes[osid] = struct{}{}
			allOses[osid] = struct{}{}
			fieldsfreq[f] = ff
		}
	}

	sort.SliceStable(fields, func(i, j int) bool {
		// Sort by number of distros using the field
		if len(fieldsfreq[fields[i]].OSes) != len(fieldsfreq[fields[j]].OSes) {
			return len(fieldsfreq[fields[i]].OSes) > len(fieldsfreq[fields[j]].OSes)
		}

		// Sort by number of files containing the field
		// if fieldsfreq[fields[i]].Count != fieldsfreq[fields[j]].Count {
		// 	return fieldsfreq[fields[i]].Count > fieldsfreq[fields[j]].Count
		// }

		// Sort by the field name
		return fieldsfreq[fields[i]].Name < fieldsfreq[fields[j]].Name
		// return fields[i] < fields[j]
	})

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)

	fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", "Count", "Field", "Spec", "Percent", "Distros")
	for _, field := range fields {
		key := fieldsfreq[field]
		official := " "
		if slices.Contains(specKeys, key.Name) {
			official = "✓"
		}
		fmt.Fprintf(writer,
			"%5d\t%s\t%4s\t%6d%%\t%s\n",
			len(key.OSes), // key.Count,
			key.Name,
			official,
			WholePercentage(len(key.OSes), len(allOses)),
			MapKeys(key.OSes),
		)
	}
	writer.Flush()
}
