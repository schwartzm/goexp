package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Take the quiz!")
	in, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
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
		fmt.Printf("user: %v, real: %v\n", user_answer, real_answer)
		if user_answer == real_answer {
			correct = correct + 1
		}
	}
	fmt.Printf("You got %v out of %v correct.\n", correct, len(ques))
}
