package tsp

import (
	"math"
	"math/rand"
	"slices"
	"time"
)

type ILS struct {
	instance         TSPInstance
	BestSolutionChan chan Solution
}

func NewILS(instance TSPInstance) ILS {
	return ILS{
		instance:         instance,
		BestSolutionChan: make(chan Solution, 10),
	}
}

func (ils *ILS) Solve(init []int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	size := len(init)
	pertStrength := int32(math.Ceil(float64(size) * 0.04))
	solution := make([]int, size)
	copy(solution, init)

	bestSolution := make([]int, size)
	copy(bestSolution, init)
	evalBestSolution := ils.instance.Evaluate(bestSolution)

	for range 100 {
		solution = ils.instance.LocalSearch(solution)
		evalSolution := ils.instance.Evaluate(solution)

		if evalSolution < evalBestSolution && math.Abs(evalSolution-evalBestSolution) > 10e-5 {
			copy(bestSolution, solution)
			evalBestSolution = evalSolution
			ils.BestSolutionChan <- Solution{Idxs: slices.Clone(bestSolution), Eval: evalBestSolution}
		}

		for range pertStrength {
			i := rng.Intn(size-1) + 1
			j := rng.Intn(size-1) + 1
			for i == j {
				j = rng.Intn(size-1) + 1
			}
			solution[i], solution[j] = solution[j], solution[i]
		}
	}
	close(ils.BestSolutionChan)
}
