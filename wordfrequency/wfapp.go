package main
/*

*/

/*
wfapp is a commandline for the wordfrequency functionality. 

Run like:

go run wfapp.go --infile someFile.txt

E.g.,

go run wfapp.go --infile wfmain.go
*/

import (
	"flag"
	"fmt"

	wf "github.com/schwartzm/goexp/wordfrequency"
)

var inFile string

func init() {
  flag.StringVar(&inFile, "infile", "", "Input file containing words") 
  flag.Parse()
}

func main(){
  fmt.Println(inFile) 
  fw,err := wf.GetWordFrequency(inFile, false, 3) 
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("Unique word count: %d\n", len(fw.Words))

  for i := 0; i < len(fw.Words); i++ {
    fmt.Printf("%s,%d\n", fw.Words[i].Word,fw.Words[i].Count)
  }
}
