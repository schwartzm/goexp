package wordfrequency

/*
Wordfrequency reports the count of each unique word in a file.
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
Wordfrequency reports the count of each unique word in a structure
 with the io.reader interface, such as a file handle from os.open().
*/

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strings"
)

type SortDir int

/*
Sort direction options. Defaults to not sorted.
*/
const (
	SORT_DESC SortDir = 0
	SORT_ASC  SortDir = 1
	SORT_NOT  SortDir = 2
)

type fileWords struct {
	File  string
	Words map[string]int
}

type Word struct {
	Word  string
	Count int
}

/*
FileWords is returned from the GetWordFrequency function.
The structure contains the path to analyzed file and
a map of the frequency of each word.
*/
type FileWords struct {
	File  string
	Words []Word
}

var keep = regexp.MustCompile("^\\w+$")

/*
GetWordFrequency returns the count of each unique word in a data source by returning
a FileWords structure. Caller will pass in one of the sort directions (e.g., wf.SORT_DESC)
, and a minimum number of characters that must be
be present in a word for it to be included.
*/
func GetWordFrequency(data io.Reader, sortDir SortDir, minWordLength int) (*FileWords, error) {

	fw := &fileWords{File: "placeholder"}
	fw.Words = make(map[string]int)

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		if len(w) >= minWordLength {
			w = strings.ToLower(strings.Trim(w, " "))
			if !keep.MatchString(w) {
				continue
			}
			fw.Words[w] = fw.Words[w] + 1
		}
	}

	return sortByFrequency(fw, sortDir), nil
}

type word struct {
	word  string
	count int
}

type byCount []word

func (c byCount) Len() int           { return len(c) }
func (c byCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCount) Less(i, j int) bool { return c[i].count < c[j].count }

/*
sortByFrequency sorts a map of counts, where key is the item being
counted and v is the count of each word. Sorts ASC or DESC.
*/
func sortByFrequency(fw *fileWords, sortDir SortDir) *FileWords {
	var words byCount

	for k, v := range fw.Words {
		words = append(words, word{k, v})
	}

	switch sortDir {
	case SORT_DESC:
		sort.Sort(sort.Reverse(byCount(words)))
	case SORT_ASC:
		sort.Sort(byCount(words))
	case SORT_NOT:
		// not sorted
	default:
		// not sorted
	}

	fwRet := FileWords{File: fw.File}

	for i := 0; i < len(words); i++ {
		fwRet.Words = append(fwRet.Words, Word{words[i].word, words[i].count})
	}

	return &fwRet
}
