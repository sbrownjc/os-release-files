package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/exp/maps"
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

// Find searches a given directory root for files and returns a list of filenames.
func Find(root string) (result []string) {
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.Type().IsRegular() {
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
func FieldsForFile(path string) (osid, version string, fields []string) {
	lines, err := FileToLines(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, l := range lines {
		split := strings.SplitN(l, "=", 2)
		key := strings.TrimSpace(split[0])
		if key != "" && !strings.HasPrefix(key, "#") {
			fields = append(fields, key)
			value := strings.Trim(strings.TrimSpace(split[1]), `"`)

			switch key {
			case "ID":
				osid = value
			case "VERSION_ID":
				version = value
			}
		}
	}

	return osid, version, fields
}

func main() {
	allOses := map[string]struct{}{}
	fileCount := 0
	fields := []string{}
	fieldsfreq := map[string]Frequency{}

	for _, s := range Find("collection") {
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

	readme, err := os.OpenFile("README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprint(readme, "# os-release-files\n\n")
	fmt.Fprint(readme, "Running `go run main.go` will update this file using the files contained in the [collection dir](./collection).\n\n")
	fmt.Fprint(readme, "Files are named after the PRETTY_NAME variable converted to lowercase and all non alphanumeric characters converted to dashes.\n\n")
	fmt.Fprint(readme, "i.e. in ZSH: `source $f; name=${PRETTY_NAME:l}; name=${name//[^a-zA-Z0-9]/-}; mv $f collection/$name`\n\n")
	fmt.Fprint(readme, "The columns in the table are:\n\n")
	fmt.Fprintln(readme, "- **COUNT**: Number of distros that contain this field")
	fmt.Fprintln(readme, "- **FIELD**: Name of this field")
	fmt.Fprintln(readme, "- **SPEC**: Is the field part of the the [os-release spec](https://www.freedesktop.org/software/systemd/man/os-release.html)?")
	fmt.Fprintln(readme, "- **PERCENT**: Percentage of distros that contain this field")
	fmt.Fprint(readme, "- **DISTROS**: List of IDs of distros that contain this field\n\n")

	table := tablewriter.NewWriter(readme)
	table.SetHeader([]string{"Count", "Field", "Spec", "Percent", "Distros"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)

	for _, field := range fields {
		key := fieldsfreq[field]
		var official string
		if slices.Contains(specKeys, key.Name) {
			official = "✓"
		}
		keys := maps.Keys(key.OSes)
		slices.Sort(keys)
		table.Append([]string{
			fmt.Sprintf("%d", len(key.OSes)),
			key.Name,
			official,
			fmt.Sprintf("%d%%", WholePercentage(len(key.OSes), len(allOses))),
			strings.Join(keys, ", "),
		})
	}
	table.Render()
}
