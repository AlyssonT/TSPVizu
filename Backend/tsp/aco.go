package tsp

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func selectRandomWeightedIndex(pheromonesWithCitiesIndexes []PheromoneWithCityIndex, normalizedWeights []float64, localRng *rand.Rand) int {
	randomValue := localRng.Float64()
	cumulative := 0.0
	for i, weight := range normalizedWeights {
		cumulative += weight
		if randomValue <= cumulative {
			return pheromonesWithCitiesIndexes[i].idx
		}
	}
	return pheromonesWithCitiesIndexes[len(pheromonesWithCitiesIndexes)-1].idx
}

type Ant struct {
	Trail []int
	Eval  float64
}

type ACO struct {
	instance   TSPInstance
	Pheromones [][]float64
	Ants       []Ant
	BestTrail  []int
	BestEval   float64
}

func NewACO(instance TSPInstance, nAnts uint) ACO {
	pheromones := make([][]float64, len(instance.Cities))
	for i := range pheromones {
		pheromones[i] = make([]float64, len(instance.Cities))
		for j := range pheromones[i] {
			pheromones[i][j] = 1.5
		}
	}
	return ACO{
		instance:   instance,
		Pheromones: pheromones,
		Ants:       make([]Ant, nAnts),
		BestTrail:  make([]int, 0, len(instance.Cities)),
		BestEval:   math.MaxInt,
	}
}

func (aco *ACO) reinforcement() {
	for i := range aco.Ants {
		aco.Ants[i].Eval = aco.instance.Evaluate(aco.Ants[i].Trail)
		for j := range len(aco.Ants[i].Trail) - 1 {
			aco.Pheromones[aco.Ants[i].Trail[j]][aco.Ants[i].Trail[j+1]] += 1.0 / float64(aco.Ants[i].Eval)
		}
	}
}

func (aco *ACO) evaporation(evaporationFactor float64) {
	for i := range aco.Pheromones {
		for j := range aco.Pheromones[i] {
			aco.Pheromones[i][j] = aco.Pheromones[i][j]*(1.0-evaporationFactor) + math.SmallestNonzeroFloat64
		}
	}
}

type PheromoneWithCityIndex struct {
	idx       int
	pheromone float64
}

func (aco *ACO) createTrails(alfa, beta, q0 float64) {
	var wg sync.WaitGroup
	wg.Add(len(aco.Ants))
	for i := range aco.Ants {
		go func(i int) {
			defer wg.Done()
			localRng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))

			visited := make([]bool, len(aco.instance.Cities))
			picked := 0
			visited[picked] = true
			aco.Ants[i].Trail = make([]int, 0, len(aco.instance.Cities))
			aco.Ants[i].Trail = append(aco.Ants[i].Trail, picked)

			for range len(aco.instance.Cities) - 1 {
				notVisitedPheromones := make([]PheromoneWithCityIndex, 0)
				for j := range len(aco.instance.Cities) {
					if !visited[j] {
						notVisitedPheromones = append(
							notVisitedPheromones,
							PheromoneWithCityIndex{idx: j, pheromone: aco.Pheromones[picked][j]},
						)
					}
				}
				if len(notVisitedPheromones) == 0 {
					break
				}

				sumProb := 0.0
				weightProbabilityCity := make([]float64, 0)
				for j := range notVisitedPheromones {
					weight := math.Pow(notVisitedPheromones[j].pheromone, alfa) * math.Pow(1.0/float64(aco.instance.Distances[picked][notVisitedPheromones[j].idx]), beta)
					weightProbabilityCity = append(weightProbabilityCity, weight)
					sumProb += weight
				}

				for j := range weightProbabilityCity {
					weightProbabilityCity[j] /= sumProb
				}

				q := localRng.Float64()
				if q < q0 {
					picked = selectRandomWeightedIndex(notVisitedPheromones, weightProbabilityCity, localRng)
				} else {
					picked = notVisitedPheromones[localRng.Intn(len(notVisitedPheromones))].idx
				}

				aco.Ants[i].Trail = append(aco.Ants[i].Trail, picked)
				visited[picked] = true
			}
		}(i)
	}
	wg.Wait()
}

func (aco *ACO) Solve(alfa, beta, q0, evaporationFactor float64, iterations, nAnts int) {
	bestEval := math.MaxFloat64
	var bestTrail []int

	for range iterations {
		aco.createTrails(alfa, beta, q0)

		iterationBestAnt := aco.Ants[0]
		for _, ant := range aco.Ants {
			if ant.Eval < iterationBestAnt.Eval {
				iterationBestAnt = ant
			}
		}

		if iterationBestAnt.Eval < bestEval {
			bestEval = iterationBestAnt.Eval
			bestTrail = make([]int, len(iterationBestAnt.Trail))
			copy(bestTrail, iterationBestAnt.Trail)
		}

		aco.reinforcement()
		aco.evaporation(evaporationFactor)
	}

	aco.BestTrail = aco.instance.LocalSearch(bestTrail)
	aco.BestEval = aco.instance.Evaluate(aco.BestTrail)
}
