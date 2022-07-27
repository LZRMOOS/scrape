// usage: go run scrape.go <directory> <keyword> <output_file>
// example: go run scrape.go /tmp/test/ /tmp/test/output

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func usage() {
	fmt.Println("usage: go run scrape.go <directory> <keyword> <output_file>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 4 {
		usage()
		os.Exit(1)
	} 

	dir := os.Args[1]
	keyword := os.Args[2]
	outfile := os.Args[3]

	// regexp object to match the keyword
	re := regexp.MustCompile(keyword)

	// the file to write to
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	w := bufio.NewWriter(f)

	// walk the directory and grep for the keyword in each file
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			data, err := ioutil.ReadFile(path)
			if err == nil {
				// split the file into lines
				lines := strings.Split(string(data), "\n")

				for _, line := range lines {
					if re.MatchString(line) {
						// strip the match and spaces from the line
						line = strings.TrimSpace(re.ReplaceAllString(line, ""))
						
						// write the line to the file
						fmt.Fprintln(w, line)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	w.Flush()
}
