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
	"time"
)

var csvPath string
var maxTime int

func init() {
	flag.StringVar(&csvPath, "config", "problems.csv", "Path to CSV configuration file")
	flag.IntVar(&maxTime, "maxtime", 15, "Maximum quiz time in seconds")
	flag.Parse()
}

func timer(max int, c chan int) {
	time.Sleep(time.Duration(max) * time.Second)
	c <- max
}

func main() {
	fmt.Println("Using config file ", csvPath)
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
	c := make(chan int)
	go timer(maxTime, c)
	correct := 0
	for _, v := range ques {
		select {
		case t := <-c: // t is assinged the value received in the channel c.
			fmt.Println("\nTimes up! ", t, " seconds")
			return
		default:
			// proceed as normal
		}

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
