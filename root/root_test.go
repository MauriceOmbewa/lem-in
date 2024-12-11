package root

import (
	"os"
	"testing"
)

func TestPrintOutPut(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Invalid file extension",
			args:     []string{"program", "test.csv"},
			expected: "only accept files in txt: .csv not allowed.\n",
		},
		{
			name:     "Valid file extension and successful simulation",
			args:     []string{"program", "test.txt"},
			expected: "\n[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()
			os.Args = tt.args

			PrintOutPut()
		})
	}
}
