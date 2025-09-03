package tsp

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type City struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Solution struct {
	Idxs []int   `json:"idxs"`
	Eval float64 `json:"eval"`
}

func DistanceBetweenCities(c1 *City, c2 *City) float64 {
	dx := c1.X - c2.X
	dy := c1.Y - c2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func reverseSlice(arr []int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type TSPInstance struct {
	Distances [][]float64
	Cities    []City
}

func NewTSPInstance(cities []City) TSPInstance {
	distances := make([][]float64, len(cities))
	for i := range distances {
		distances[i] = make([]float64, len(cities))
	}
	for i, c1 := range cities {
		for j := i + 1; j < len(cities); j++ {
			dist := DistanceBetweenCities(&c1, &cities[j])
			distances[i][j] = dist
			distances[j][i] = dist
		}
	}

	return TSPInstance{
		Cities:    cities,
		Distances: distances,
	}
}

func NewTSPInstanceFromFile(fileName string) TSPInstance {
	file, err := os.Open("datasets/" + fileName)
	if err != nil {
		log.Fatal("Error on opening dataset file")
	}

	reader := bufio.NewReader(file)
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error on reading dataset file")
	}

	nCities, err := strconv.ParseInt(strings.Fields(firstLine)[0], 10, 32)
	if err != nil {
		log.Fatal("Error on parsing the number of cities")
	}

	cities := make([]City, nCities)
	for {
		rawLine, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		values := strings.Fields(rawLine)
		if len(values) < 3 {
			log.Fatal("Invalid file")
		}
		idx, err1 := strconv.ParseInt(values[0], 10, 32)
		x, err2 := strconv.ParseFloat(values[1], 64)
		y, err3 := strconv.ParseFloat(values[2], 64)

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatal("Invalid file")
		}

		cities[idx-1] = City{X: x, Y: y}
	}

	return NewTSPInstance(cities)
}

func (instance *TSPInstance) Evaluate(solution []int) float64 {
	result := 0.0
	n := len(solution)
	for i := range n {
		next := (i + 1) % n
		result += instance.Distances[solution[i]][solution[next]]
	}
	return result
}

func (s *TSPInstance) LocalSearch(init []int) []int {
	size := len(init)
	currentSolution := make([]int, size)
	copy(currentSolution, init)
	currentEval := s.Evaluate(currentSolution)

	for {
		bestSolution := make([]int, size)
		bestEval := math.MaxFloat64

		for i := 1; i < size-1; i++ {
			for j := i + 1; j < size; j++ {
				tempEval := currentEval
				tempEval -= s.Distances[currentSolution[i-1]][currentSolution[i]]
				tempEval += s.Distances[currentSolution[i-1]][currentSolution[j]]
				if j == size-1 {
					tempEval -= s.Distances[currentSolution[j]][currentSolution[0]]
					tempEval += s.Distances[currentSolution[i]][currentSolution[0]]
				} else {
					tempEval -= s.Distances[currentSolution[j]][currentSolution[j+1]]
					tempEval += s.Distances[currentSolution[i]][currentSolution[j+1]]
				}

				if tempEval < bestEval {
					bestEval = tempEval
					copy(bestSolution, currentSolution)
					reverseSlice(bestSolution, i, j)
				}
			}
		}

		if bestEval < currentEval {
			currentEval = bestEval
			copy(currentSolution, bestSolution)
		} else {
			return currentSolution
		}
	}
}
