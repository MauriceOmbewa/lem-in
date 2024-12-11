package root

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"lem-in/utils"
)

func PrintOutPut() {
	args := os.Args[1]
	path := filepath.Ext(args)
	if path != ".txt" {
		fmt.Printf("only accept files in txt: %s not allowed.\n", path)
		return
	}

	antPop := &utils.AntPopulation{}
	endPoints := &utils.EndPoints{}

	graph, _, err := utils.BuildGraph(args, antPop, endPoints)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	start := endPoints.Start
	end := endPoints.End

	paths := utils.FindAllPaths(graph, start, end)

	// Sort the 2D slice by the length of each sub-slice
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	unique := utils.UniquePaths(paths)

	antAssignments := utils.AssignPathsToAnts(antPop.Size, unique)

	// Create AntFarm using antAssignments
	farm := utils.CreateAntFarm(antAssignments, start, end)

	// Simulate movement
	moves, err := farm.SimulateMovement()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	utils.Readfile(args)

	fmt.Println()
	fmt.Print(moves)
}
