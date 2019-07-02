package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var flagPath string

func init() {
	flag.StringVar(&flagPath, "config", "problems.csv", "Path to CSV configuration file")
	flag.Parse()
}

func main() {
	fmt.Println("Using config file ", flagPath)
	in, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	r := csv.NewReader(in)
	ques, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	correct := 0
	for _, v := range ques {
		fmt.Printf("What is %v? ", v[0])
		ans, _ := reader.ReadString('\n')
		ans = strings.Replace(ans, "\n", "", -1)
		user_answer, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println(err)
			continue
		}
		real_answer, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal("Cannot interpret CSV value %v for question %v.", v[1], v[0])
		}
		if user_answer == real_answer {
			correct = correct + 1
		}
	}
	fmt.Printf("You got %v out of %v correct.\n", correct, len(ques))
}
