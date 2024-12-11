package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// FindAllPaths uses Depth First Search (dfs) algorithm to recursively
// check for all paths that are available in the graph
func FindAllPaths(graph Graph, start, end string) [][]string {
	var paths [][]string
	var dfs func(string, []string)

	dfs = func(current string, path []string) {
		path = append(path, current)

		if current == end {
			paths = append(paths, append([]string(nil), path...))
			return
		}

		for _, neighbor := range graph[current] {
			if !contains(path, neighbor) {
				dfs(neighbor, path)
			}
		}
	}

	dfs(start, []string{})
	return paths
}

func Readfile(file string) {
	fileName, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("could not red file: %s\n", err)
		return
	}
	fmt.Println(string(fileName))
}

// UniquePaths returns the set of valid paths that do not share any room and
// provides the path with the most number of paths for optimum and faster
// movement of ants
func UniquePaths(paths [][]string) [][]string {
	var allUniquePaths [][][]string

	for i, basePath := range paths {
		unique := make([][]string, 0)
		unique = append(unique, basePath)
		m1 := make(map[string]bool)

		// Mark all rooms in the current base path as visited
		for _, val := range basePath {
			m1[val] = true
		}

		// Check for unique paths relative to the current base path
		for j, path := range paths {
			if i == j {
				continue // Skip the current base path itself
			}

			slice := path[1 : len(path)-1] // Ignore start and end rooms for uniqueness
			isUnique := func([]string) bool {
				for _, val := range slice {
					if _, ok := m1[val]; ok {
						return false
					}
				}
				return true
			}

			if isUnique(path) {
				unique = append(unique, path)
				// Mark the rooms in this unique path as visited
				for _, value := range slice {
					m1[value] = true
				}
			}
		}

		allUniquePaths = append(allUniquePaths, unique)
	}

	return optimumPaths(allUniquePaths)
}

// optimumPaths Checks for the set of valid paths that has more
// paths than any other set
func optimumPaths(slices [][][]string) [][]string {
	var longest [][]string

	for _, twoDSlice := range slices {
		if len(twoDSlice) > len(longest) {
			longest = twoDSlice
		}
	}

	return longest
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func AssignPathsToAnts(antCount int, paths [][]string) map[int][]string {
	assignments := make(map[int][]string)
	pathAnts := make([]int, len(paths))

	for ant := 1; ant <= antCount; ant++ {
		assignedPath := 0
		for i := 1; i < len(paths); i++ {
			if len(paths[i-1])+pathAnts[i-1] >= len(paths[i])+pathAnts[i] {
				assignedPath = i
			} else {
				break
			}
		}
		assignments[ant] = paths[assignedPath]
		pathAnts[assignedPath]++
	}

	return assignments
}

func CreateAntFarm(antAssignments map[int][]string, start, end string) *AntFarm {
	farm := &AntFarm{
		Ants:  make([]*Ant, len(antAssignments)),
		Rooms: make(map[string]*Room),
	}

	// Create rooms
	for _, path := range antAssignments {
		for _, roomName := range path {
			if _, exists := farm.Rooms[roomName]; !exists {
				room := &Room{Name: roomName}
				if roomName == start {
					room.IsStart = true
					farm.Start = room
				} else if roomName == end {
					room.IsEnd = true
					farm.End = room
				}
				farm.Rooms[roomName] = room
			}
		}
	}

	// Create ants and assign paths
	for antID, pathRooms := range antAssignments {
		ant := &Ant{
			Id:          antID,
			Path:        make([]*Room, len(pathRooms)),
			PathIndex:   0,
			CurrentRoom: farm.Start,
			HasReached:  false,
		}
		for j, roomName := range pathRooms {
			ant.Path[j] = farm.Rooms[roomName]
		}
		farm.Ants[antID-1] = ant
	}

	return farm
}

// Simulating ants movements.
func (af *AntFarm) SimulateMovement() (string, error) {
	if len(af.Ants) == 0 {
		return "", errors.New("invalid data format, invalid number of Ants")
	}

	occupiedRooms := make(map[*Room]*Ant)
	var allMoves []string

	for !af.areAllAntsReached() {
		moves := af.performAntMoves(occupiedRooms)
		if len(moves) > 0 {
			allMoves = append(allMoves, strings.Join(moves, " "))
		}
		if af.Move != 0 {
			af.Move = 0
		}
	}

	return strings.Join(allMoves, "\n") + "\n", nil
}

func (af *AntFarm) areAllAntsReached() bool {
	for _, ant := range af.Ants {
		if err := af.validateAnt(ant); err != nil {
			return true
		}
		if !ant.HasReached {
			return false
		}
	}
	return true
}

func (af *AntFarm) validateAnt(ant *Ant) error {
	if ant == nil {
		return errors.New("ant is nil")
	}
	if len(ant.Path) == 0 {
		return fmt.Errorf("ant %d has no valid path", ant.Id)
	}
	if ant.CurrentRoom == nil {
		return fmt.Errorf("ant %d has no current room set", ant.Id)
	}
	if len(ant.Path) == 2 && af.Move != 0 {
		return fmt.Errorf("only move one ant per turn")

	}
	return nil
}

func (af *AntFarm) performAntMoves(occupiedRooms map[*Room]*Ant) []string {
	var moves []string

	// Clear previously occupied non-start/end rooms
	for room := range occupiedRooms {
		if !room.IsStart && !room.IsEnd {
			occupiedRooms[room] = nil
		}
	}

	for _, ant := range af.Ants {
		if err := af.validateAnt(ant); err != nil {
			continue
		}

		if ant.HasReached {
			continue
		}

		move := af.moveAnt(ant, occupiedRooms)
		if len(ant.Path) == 2 {
			af.Move++
		}
		if move != "" {
			moves = append(moves, move)
		}
	}

	return moves
}

func (af *AntFarm) moveAnt(ant *Ant, occupiedRooms map[*Room]*Ant) string {
	if ant.PathIndex >= len(ant.Path)-1 {
		return ""
	}

	nextRoom := ant.Path[ant.PathIndex+1]

	// Check if next room is available or is the end room
	if occupiedRooms[nextRoom] != nil && !nextRoom.IsEnd {
		return ""
	}

	// Clear previous room if not start/end
	if !ant.CurrentRoom.IsStart && !ant.CurrentRoom.IsEnd {
		occupiedRooms[ant.CurrentRoom] = nil
	}

	// Move ant to next room
	ant.CurrentRoom = nextRoom
	ant.PathIndex++

	// Mark room as occupied if not start/end
	if !nextRoom.IsStart && !nextRoom.IsEnd {
		occupiedRooms[nextRoom] = ant
	}

	// Check if ant has reached end
	if nextRoom.IsEnd {
		ant.HasReached = true
	}

	return fmt.Sprintf("L%d-%s", ant.Id, nextRoom.Name)
}
