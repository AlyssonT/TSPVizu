package handlers

import (
	"time"

	"github.com/AlyssonT/tsp-visual-backend/tsp"
	"github.com/gin-gonic/gin"
)

type SolveRequest struct {
	Cities   []tsp.City `json:"cities"`
	FileName string     `json:"fileName"`
}

func Range[T any](slice []T) []int {
	result := make([]int, len(slice))
	for i := range slice {
		result[i] = i
	}
	return result
}

func TSP_Solve(ctx *gin.Context) {
	var request SolveRequest
	ctx.ShouldBindJSON(&request)

	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	clientGone := ctx.Writer.CloseNotify()

	var instance tsp.TSPInstance
	if len(request.Cities) > 0 {
		instance = tsp.NewTSPInstance(request.Cities)
	} else if request.FileName != "" {
		instance = tsp.NewTSPInstanceFromFile(request.FileName)
	} else {
		instance = tsp.NewTSPInstanceFromFile("dsj1000.tsp")
	}

	sequential := Range(instance.Cities)
	ils := tsp.NewILS(instance)
	go ils.Solve(sequential)

	for solution := range ils.BestSolutionChan {
		select {
		case <-clientGone:
			return
		default:
			ctx.SSEvent("message", solution)
			ctx.Writer.Flush()

			time.Sleep(1000 * time.Millisecond)
		}
	}

	ctx.Writer.Flush()
}
