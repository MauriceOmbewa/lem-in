package utils

import (
	"reflect"
	"testing"
)

func TestFindAllPaths(t *testing.T) {
	type args struct {
		graph Graph
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Single Path",
			args: args{
				graph: Graph{
					"A": {"B"},
					"B": {"C"},
					"C": {},
				},
				start: "A",
				end:   "C",
			},
			want: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name: "Multiple Paths",
			args: args{
				graph: Graph{
					"A": {"B", "C"},
					"B": {"D"},
					"C": {"D"},
					"D": {},
				},
				start: "A",
				end:   "D",
			},
			want: [][]string{
				{"A", "B", "D"},
				{"A", "C", "D"},
			},
		},
		{
			name: "No Path",
			args: args{
				graph: Graph{
					"A": {"B"},
					"B": {"C"},
					"C": {},
				},
				start: "A",
				end:   "D",
			},
			want: [][]string{},
		},
		{
			name: "Cycle in Graph",
			args: args{
				graph: Graph{
					"A": {"B"},
					"B": {"C", "A"},
					"C": {"D"},
					"D": {},
				},
				start: "A",
				end:   "D",
			},
			want: [][]string{
				{"A", "B", "C", "D"},
			},
		},
		{
			name: "Start Equals End",
			args: args{
				graph: Graph{
					"A": {"B"},
					"B": {"C"},
					"C": {"A"},
				},
				start: "A",
				end:   "A",
			},
			want: [][]string{
				{"A"},
			},
		},
		{
			name: "Disconnected Graph",
			args: args{
				graph: Graph{
					"A": {"B"},
					"C": {"D"},
				},
				start: "A",
				end:   "D",
			},
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAllPaths(tt.args.graph, tt.args.start, tt.args.end)

			// Debugging: Print out the found paths for inspection
			t.Logf("Test: %s - Found Paths: %v", tt.name, got)

			// Check if the lengths match
			if len(got) != len(tt.want) {
				t.Errorf("FindAllPaths() = %v, want %v", got, tt.want)
			} else {
				// If lengths are equal, check element by element
				for i := range got {
					if !reflect.DeepEqual(got[i], tt.want[i]) {
						t.Errorf("FindAllPaths() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
