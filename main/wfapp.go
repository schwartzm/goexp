package main

/*
wfapp.go is the CLI program for wordfrequency.go.
Copyright (C) 2018  Michael Schwartz

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
wfapp is a commandline interface to wordfrequency.go.

Run like:

./wfapp --infile input.txt [--sort a|d|n]

E.g.,

./wfapp --infile wfmain.go --sort a

.wfapp --help

*/

import (
	"flag"
	"fmt"
	"os"

	wf "github.com/schwartzm/goexp/wordfrequency"
)

var inFile, sortDirFlag string

func init() {
	const (
		infileHelp  = "Path to file to process (required)."
		defaultSort = "d"
		sortHelp    = "Sort results. Desc (d), Asc (a), Not (n or unset)."
	)
	flag.StringVar(&inFile, "infile", "", infileHelp)
	flag.StringVar(&sortDirFlag, "sort", defaultSort, sortHelp)
	flag.Parse()
}

func main() {
	fmt.Println(inFile)

	file, err := os.Open(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	sortDir := wf.SORT_NOT

	switch sortDirFlag {
	case "a":
		sortDir = wf.SORT_ASC
	case "d":
		sortDir = wf.SORT_DESC
	case "n":
		sortDir = wf.SORT_NOT
	default:
		sortDir = wf.SORT_NOT

	}

	fw, err := wf.GetWordFrequency(file, sortDir, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Unique word count: %d\n", len(fw.Words))

	for i := 0; i < len(fw.Words); i++ {
		fmt.Printf("%s,%d\n", fw.Words[i].Word, fw.Words[i].Count)
	}
}
