package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type job struct {
	weight int
	length int
	score  float64
}

type byScore []job

func (a byScore) Len() int      { return len(a) }
func (a byScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool {
	if a[i].score == a[j].score {
		return a[i].weight > a[j].weight
	}
	return a[i].score > a[j].score
}

func parseJobs(fname string) ([]job, error) {
	content, readErr := ioutil.ReadFile(fname)
	if readErr != nil {
		return nil, readErr
	}

	var jobs []job

	for index, line := range strings.Split(string(content), "\n") {

		// First line contains the number of jobs
		if index == 0 {
			size, sizeErr := strconv.ParseInt(line, 10, 32)
			if sizeErr != nil {
				return nil, sizeErr
			}
			jobs = make([]job, size)
			continue
		}

		// A line contain weight and length of a job separated by whitespace
		splitted := strings.Fields(line)
		if len(splitted) == 0 {
			break
		}

		weightStr, lengthStr := splitted[0], splitted[1]

		weight, weightParseErr := strconv.ParseInt(weightStr, 10, 32)
		if weightParseErr != nil {
			return nil, weightParseErr
		}

		length, lengthParseErr := strconv.ParseInt(lengthStr, 10, 32)
		if lengthParseErr != nil {
			return nil, lengthParseErr
		}

		jobs[index-1] = job{
			weight: int(weight),
			length: int(length),
			// score:  float64(weight) - float64(length)}  // problem 1
			score: float64(weight) / float64(length)} // problem 2

	}

	return jobs, nil
}

func main() {
	jobs, parseErr := parseJobs("./jobs.txt")
	if parseErr != nil {
		log.Fatal("Could not parse the data", parseErr)
	}
	sort.Sort(byScore(jobs))

	completionTime := 0
	weightedSum := 0
	for _, job := range jobs {
		completionTime += job.length
		weightedSum += completionTime * job.weight
	}

	// Right answer is 69119377652
	fmt.Println(weightedSum)
}
