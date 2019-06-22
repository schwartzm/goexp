package wordfrequency

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
)

type fileWords struct {
	File  string
	Words map[string]int
}

type Word struct {
  Word string
  Count int
}

type FileWords struct {
  File string
  Words []Word 
}

var keep = regexp.MustCompile("^\\w+$")

func GetWordFrequency(inFile string, sortDesc bool, minWordLength int) (*FileWords, error) {
	file, err := os.Open(inFile)
	if err != nil {
		return nil, err
	}

	fw := &fileWords{File: inFile}
	fw.Words = make(map[string]int)

	defer file.Close()

	scanner := bufio.NewScanner(file)
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

	return sortByFrequency(fw, "DESC"), nil
}

type word struct {
  word string
  count int
}

type byCount []word

func (c byCount) Len() int {return len(c)}
func (c byCount) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c byCount) Less(i, j int) bool { return c[i].count < c[j].count }

func sortByFrequency(fw *fileWords, sortDir string)(*FileWords) {
  var words byCount

  for k,v := range fw.Words {
    words = append(words, word{ k, v })
  }

  switch sortDir {
    case "DESC":
      sort.Sort(sort.Reverse(byCount(words)))
    case "ASC":
      sort.Sort(byCount(words))
    default:
      // not sorted
  }

  fwRet := FileWords{ File: fw.File }
  
  for i := 0; i < len(words); i++ {
   fwRet.Words = append(fwRet.Words, Word{ words[i].word, words[i].count })
  }

  return &fwRet
   /* 
  for k := range fw.Words {
    delete(fw.Words,k)
  }
  for i := 0; i < len(words); i++ {
    fw.Words[words[i].word] = words[i].count       
  }

  return fw
  */
}

