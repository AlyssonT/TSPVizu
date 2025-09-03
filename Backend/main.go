package main

import (
	"github.com/AlyssonT/tsp-visual-backend/configs"
	"github.com/AlyssonT/tsp-visual-backend/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	configs.BuildConfigs()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{configs.GetConfigs().FrontendURL}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	server.Use(cors.New((config)))

	server.POST("/solve", handlers.TSP_Solve)

	server.Run()

	// instance := tsp.NewTSPInstanceFromFile("st70.tsp")

	// sequential := handlers.Range(instance.Cities)
	// ils := tsp.NewILS(instance)
	// bestEval := math.MaxInt
	// bestSolution := make([]int, len(instance.Cities))

	// start := time.Now()
	// type Result struct {
	// 	Solution []int
	// 	Eval     int
	// }

	// resultChan := make(chan Result, 2500)

	// for range 2500 {
	// 	go func() {
	// 		solution, eval := ils.Solve(sequential)
	// 		resultChan <- Result{Solution: solution, Eval: eval}
	// 	}()
	// }

	// for range 2500 {
	// 	result := <-resultChan
	// 	if result.Eval < bestEval {
	// 		bestEval = result.Eval
	// 		bestSolution = result.Solution
	// 	}
	// }

	// elapsed := time.Since(start)

	// fmt.Println(elapsed)
	// fmt.Println(bestSolution, bestEval)

	// aco := tsp.NewACO(instance, 10)
	// start = time.Now()
	// aco.Solve(1.0, 5.0, 0.96, 0.1, 30000, 100)
	// elapsed = time.Since(start)

	// fmt.Println(elapsed)
	// fmt.Println(aco.BestTrail, aco.BestEval)
}
