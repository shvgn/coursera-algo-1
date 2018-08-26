package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
In this programming problem you'll code up Prim's minimum spanning tree algorithm.

Download the text file below (right click on the file and select "Save As..."): edges.txt

This file describes an undirected graph with integer edge costs. It has the format

[number_of_nodes] [number_of_edges]

[one_node_of_edge_1] [other_node_of_edge_1] [edge_1_cost]

[one_node_of_edge_2] [other_node_of_edge_2] [edge_2_cost]

...

For example, the third line of the file is "2 3 -8874", indicating that there is an edge connecting vertex #2 and vertex
#3 that has cost -8874.

You should NOT assume that edge costs are positive, nor should you assume that they are distinct.

Your task is to run Prim's minimum spanning tree algorithm on this graph.

You should report the overall cost of a minimum spanning tree --- an integer, which may or may not be negative --- in
the box below.

IMPLEMENTATION NOTES: This graph is small enough that the straightforward O(mn) time implementation of Prim's algorithm
should work fine. OPTIONAL: For those of you seeking an additional challenge, try implementing a heap-based version. The
simpler approach, which should already give you a healthy speed-up, is to maintain relevant edges in a heap (with keys =
edge costs). The superior approach stores the unprocessed vertices in the heap, as described in lecture. Note this
requires a heap that supports deletions, and you'll probably need to maintain some kind of mapping between vertices and
their positions in the heap.
*/

type node struct {
	index           int
	edges           []*edge
	minimalEdgeCost int
}

func (n *node) resetMinEdgeCost() {
	minCost := math.MaxInt32
	for _, edge := range n.edges {
		if edge.cost < minCost {
			minCost = edge.cost
		}
	}
}

type edge struct {
	node1     int
	node2     int
	cost      int
	accounted bool
}

func parseCount(s *bufio.Scanner) (int, int, error) {
	ok := s.Scan()
	if !ok {
		return 0, 0, s.Err()
	}
	line := s.Text()
	return parseLine(line)

}

func parseLine(line string) (int, int, error) {
	var err error

	splitted := strings.Fields(line)
	if len(splitted) != 2 {
		err = fmt.Errorf("unexpected line: \"%s\"", line)
		return 0, 0, err
	}

	nodesCountStr, edgesCountStr := splitted[0], splitted[1]

	nodesCount, nodesCountParseErr := strconv.ParseInt(nodesCountStr, 10, 32)
	if nodesCountParseErr != nil {
		err = fmt.Errorf("cannot parse nodes count: %s", nodesCountParseErr)
		return 0, 0, err
	}

	edgesCount, edgesCountParseErr := strconv.ParseInt(edgesCountStr, 10, 32)
	if edgesCountParseErr != nil {
		err = fmt.Errorf("cannot parse edges count: %s", edgesCountParseErr)
		return 0, 0, err
	}

	return int(nodesCount), int(edgesCount), nil
}

func main() {

	filenmae := "./edges.txt"
	file, openErr := os.Open(filenmae)
	if openErr != nil {
		log.Fatal("Cannot open file", openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	Nnodes, Nedges, countParseErr := parseCount(scanner)
	if countParseErr != nil {
		log.Fatal(countParseErr)
	}

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
