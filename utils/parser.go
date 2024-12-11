package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[string][]string

type AntPopulation struct {
	Size int
}

type EndPoints struct {
	Start string
	End   string
}

func BuildGraph(roomFile string, antPopulation *AntPopulation, endPoints *EndPoints) (Graph, map[string][]int, error) {
	// Open the file
	file, err := os.Open(roomFile)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	// Initialize data structures
	graph := make(Graph)
	rooms := make(map[string][]int)

	scanner := bufio.NewScanner(file)

	var isStart, isEnd bool

	// Check if the file is empty (no lines at all)
	if !scanner.Scan() {
		return nil, nil, errors.New("invalid data format, invalid number of Ants") // Treat empty file as invalid ant population error
	}

	// Get ant population
	text := strings.TrimSpace(scanner.Text())
	if text == "" {
		return nil, nil, errors.New("invalid data format, invalid number of Ants") // Treat first line as invalid if it's empty
	}

	// Expecting ant population as a single integer
	if size, err := strconv.Atoi(text); err == nil {
		antPopulation.Size = size
	} else {
		return nil, nil, errors.New("invalid data format, invalid number of Ants") // Invalid ant population format
	}

	// Parse rooms and connections
	for scanner.Scan() {
		text = strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		// Detect start and end markers
		if text == "##start" {
			isStart = true
			continue
		}
		if text == "##end" {
			isEnd = true
			continue
		}

		if strings.HasPrefix(text, "#") {
			continue
		}

		// Handle the line after start/end markers
		if isStart {
			endPoints.Start = strings.Fields(text)[0]
			isStart = false
		}
		if isEnd {
			endPoints.End = strings.Fields(text)[0]
			isEnd = false
		}

		// Handle room data (e.g., "RoomName X Y")
		parts := strings.Fields(text)
		if len(parts) == 3 {
			coords := make([]int, 2)
			coords[0], err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, errors.New("invalid data format, invalid room X coordinate")
			}
			coords[1], err = strconv.Atoi(parts[2])
			if err != nil {
				return nil, nil, errors.New("invalid data format, invalid room Y coordinate")
			}
			rooms[parts[0]] = coords
			continue
		}

		// Handle connections (e.g., "Room1-Room2")
		if isTunnel(parts) {
			connection := strings.Split(parts[0], "-")
			if len(connection) != 2 {
				return nil, nil, errors.New("invalid data format, invalid tunnel")
			}

			if connection[0] == connection[1] {
				return nil, nil, errors.New("invalid data format, tunnel connects to itself")
			}

			graph[connection[0]] = append(graph[connection[0]], connection[1])
			graph[connection[1]] = append(graph[connection[1]], connection[0])
			continue
		}

		return nil, nil, errors.New("unexpected line format")
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return graph, rooms, nil
}

func isTunnel(parts []string) bool {
	return len(parts) == 1 && strings.Contains(parts[0], "-")
}
