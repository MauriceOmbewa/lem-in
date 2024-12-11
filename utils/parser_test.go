// package utils

// import (
// 	"reflect"
// 	"testing"
// )

// func TestBuildGraph(t *testing.T) {
// 	type args struct {
// 		roomFile      string
// 		antPopulation *AntPopulation
// 		endPoints     *EndPoints
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    Graph
// 		want1   map[string][]int
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, err := BuildGraph(tt.args.roomFile, tt.args.antPopulation, tt.args.endPoints)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("BuildGraph() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("BuildGraph() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("BuildGraph() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	// Helper function to create a temporary test file
	createTempFile := func(content string) (string, error) {
		tempDir := t.TempDir()
		tempFile := filepath.Join(tempDir, "test_input.txt")
		err := os.WriteFile(tempFile, []byte(content), 0644)
		return tempFile, err
	}

	// Helper function to deep compare graphs
	compareGraphs := func(expected, actual Graph) bool {
		if len(expected) != len(actual) {
			return false
		}

		for key, expectedAdj := range expected {
			actualAdj, exists := actual[key]
			if !exists {
				return false
			}

			if len(expectedAdj) != len(actualAdj) {
				return false
			}

			// Create maps to compare without order dependency
			expectedMap := make(map[string]bool)
			actualMap := make(map[string]bool)

			for _, v := range expectedAdj {
				expectedMap[v] = true
			}

			for _, v := range actualAdj {
				actualMap[v] = true
			}

			for k := range expectedMap {
				if !actualMap[k] {
					return false
				}
			}
		}
		return true
	}

	// Helper function to deep compare room coordinates
	compareRooms := func(expected, actual map[string][]int) bool {
		if len(expected) != len(actual) {
			return false
		}

		for key, expectedCoords := range expected {
			actualCoords, exists := actual[key]
			if !exists {
				return false
			}

			if len(expectedCoords) != len(actualCoords) {
				return false
			}

			for i, coord := range expectedCoords {
				if coord != actualCoords[i] {
					return false
				}
			}
		}
		return true
	}

	tests := []struct {
		name          string
		fileContent   string
		expectedGraph Graph
		expectedRooms map[string][]int
		expectedAnts  int
		expectedStart string
		expectedEnd   string
		expectedError bool
		errorMessage  string
	}{
		{
			name: "Valid Input with Rooms and Tunnels",
			fileContent: `10
Room1 0 0
Room2 1 1
Room3 2 2
##start
StartRoom 3 3
##end
EndRoom 4 4
Room1-Room2
Room2-Room3
StartRoom-Room1
EndRoom-Room3`,
			expectedGraph: Graph{
				"Room1":     []string{"Room2", "StartRoom"},
				"Room2":     []string{"Room1", "Room3"},
				"Room3":     []string{"Room2", "EndRoom"},
				"StartRoom": []string{"Room1"},
				"EndRoom":   []string{"Room3"},
			},
			expectedRooms: map[string][]int{
				"Room1":     {0, 0},
				"Room2":     {1, 1},
				"Room3":     {2, 2},
				"StartRoom": {3, 3},
				"EndRoom":   {4, 4},
			},
			expectedAnts:  10,
			expectedStart: "StartRoom",
			expectedEnd:   "EndRoom",
			expectedError: false,
		},
		{
			name: "Invalid Ant Population",
			fileContent: `ten
Room1 0 0
Room2 1 1`,
			expectedError: true,
			errorMessage:  "invalid data format, invalid number of Ants",
		},
		{
			name: "Self-Connecting Tunnel",
			fileContent: `5
Room1 0 0
Room1-Room1`,
			expectedError: true,
			errorMessage:  "invalid data format, tunnel connects to itself",
		},
		{
			name: "Invalid Room Coordinates",
			fileContent: `5
Room1 zero 0
Room2 1 1`,
			expectedError: true,
			errorMessage:  "invalid data format, invalid room X coordinate",
		},
		{
			name:          "Empty File",
			fileContent:   ``,
			expectedError: true,
			errorMessage:  "invalid data format, invalid number of Ants",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary test file
			tempFile, err := createTempFile(tt.fileContent)
			if err != nil {
				t.Fatalf("Error creating temporary test file: %v", err)
			}

			// Initialize structs
			antPopulation := &AntPopulation{}
			endPoints := &EndPoints{}

			// Call the function
			graph, rooms, err := BuildGraph(tempFile, antPopulation, endPoints)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
					return
				}

				// Check if error message contains expected message
				if !strings.Contains(err.Error(), tt.errorMessage) {
					t.Errorf("Expected error message containing '%s', got '%v'", tt.errorMessage, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				// Validate Graph
				if !compareGraphs(tt.expectedGraph, graph) {
					t.Errorf("Graph does not match expected graph\nExpected: %v\nGot: %v", tt.expectedGraph, graph)
				}

				// Validate Rooms
				if !compareRooms(tt.expectedRooms, rooms) {
					t.Errorf("Rooms do not match expected rooms\nExpected: %v\nGot: %v", tt.expectedRooms, rooms)
				}

				// Validate Ant Population
				if antPopulation.Size != tt.expectedAnts {
					t.Errorf("Ant population does not match\nExpected: %d\nGot: %d", tt.expectedAnts, antPopulation.Size)
				}

				// Validate Start and End Points
				if endPoints.Start != tt.expectedStart {
					t.Errorf("Start point does not match\nExpected: %s\nGot: %s", tt.expectedStart, endPoints.Start)
				}
				if endPoints.End != tt.expectedEnd {
					t.Errorf("End point does not match\nExpected: %s\nGot: %s", tt.expectedEnd, endPoints.End)
				}
			}
		})
	}
}

func TestBuildGraphFileError(t *testing.T) {
	// Test non-existent file
	antPopulation := &AntPopulation{}
	endPoints := &EndPoints{}

	_, _, err := BuildGraph("/path/to/non/existent/file", antPopulation, endPoints)
	if err == nil {
		t.Error("Expected error when file does not exist, but got none")
	} else if !strings.Contains(err.Error(), "could not open file") {
		t.Errorf("Unexpected error message: %v", err)
	}
}
