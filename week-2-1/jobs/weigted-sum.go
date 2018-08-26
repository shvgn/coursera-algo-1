package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type job struct {
	weight int
	length int
	score  float64
}

// Sorting
type byScore []job

func (a byScore) Len() int {
	return len(a)

}
func (a byScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byScore) Less(i, j int) bool {
	if a[i].score == a[j].score {
		return a[i].weight > a[j].weight
	}
	return a[i].score > a[j].score
}

// Parsing

func parseJobCount(s *bufio.Scanner) (int, error) {
	ok := s.Scan()
	if !ok {
		log.Fatal(s.Err())
	}
	line := s.Text()
	count, err := strconv.ParseInt(line, 10, 32)
	return int(count), err
}

func parseLine(line string) (int, int, error) {
	var err error

	splitted := strings.Fields(line)
	if len(splitted) != 2 {
		err = fmt.Errorf("unexpected line: \"%s\"", line)
		return 0, 0, err
	}

	weightStr, lengthStr := splitted[0], splitted[1]

	weight, weightParseErr := strconv.ParseInt(weightStr, 10, 32)
	if weightParseErr != nil {
		err = fmt.Errorf("cannot parse job weight: %s", weightParseErr)
		return 0, 0, err
	}

	length, lengthParseErr := strconv.ParseInt(lengthStr, 10, 32)
	if lengthParseErr != nil {
		err = fmt.Errorf("cannot parse job weight: %s", lengthParseErr)
		return 0, 0, err
	}

	return int(weight), int(length), nil
}

func main() {

	filenmae := "./jobs.txt"
	file, openErr := os.Open(filenmae)
	if openErr != nil {
		log.Fatal("Cannot open file", openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	jobsCount, countParseErr := parseJobCount(scanner)
	if countParseErr != nil {
		log.Fatal(countParseErr)
	}

	jobs := make([]job, jobsCount)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		weight, length, err := parseLine(line)
		if err != nil {
			log.Fatal("cannot parse file", err)
		}

		// problem #1
		// score := float64(weight) - float64(length)
		// problem #2
		score := float64(weight) / float64(length)

		jobs[i] = job{
			weight: int(weight),
			length: int(length),
			score:  score,
		}

		i++
	}

	sort.Sort(byScore(jobs))

	completionTime := 0
	weightedSum := 0
	for _, job := range jobs {
		completionTime += job.length
		weightedSum += completionTime * job.weight
	}

	// Right answers are
	// #1: 69119377652
	// #2: 67311454237
	fmt.Println(weightedSum)
}
