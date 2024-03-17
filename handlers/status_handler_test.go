package handlers

import (
	"demyst-todo/types"
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestParseFlag(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected types.Config
		wantErr  bool
		errMsg   string
	}{
		{
			name: "Valid Flags",
			args: []string{"--limit=10", "--pattern=even", "--input=api", "--location=http://example.com"},
			expected: types.Config{
				Limit:    10,
				Pattern:  "even",
				Input:    "api",
				Location: "http://example.com",
			},
			wantErr: false,
			errMsg:  "",
		},
		{
			name:     "Invalid Limit",
			args:     []string{"--limit=-5", "--pattern=odd", "--input=file", "--location=/path/to/file"},
			expected: types.Config{},
			wantErr:  true,
			errMsg:   "limit cannot be 0 or negative",
		},
		{
			name:     "Invalid Pattern",
			args:     []string{"--limit=5", "--pattern=dd", "--input=file", "--location=/path/to/file"},
			expected: types.Config{},
			wantErr:  true,
			errMsg:   "pattern must be even, odd or all",
		},
		{
			name:     "Invalid Input",
			args:     []string{"--limit=5", "--pattern=odd", "--input=kafka", "--location=/path/to/file"},
			expected: types.Config{},
			wantErr:  true,
			errMsg:   "input must be api or file",
		},
		{
			name:     "Invalid Input",
			args:     []string{"--limit=5", "--pattern=odd", "--input=api", "--location="},
			expected: types.Config{},
			wantErr:  true,
			errMsg:   "location cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := cli.NewApp()
			set := flag.NewFlagSet(tt.name, flag.ExitOnError)
			set.String("limit", "", "")
			set.String("pattern", "", "")
			set.String("input", "", "")
			set.String("location", "", "")
			set.Parse(tt.args)
			ctx := cli.NewContext(app, set, nil)

			got, err := parseFlag(ctx)
			fmt.Println("ddd", got, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFlag() error = %v, wantErr %v", err, tt.errMsg)
				return
			}

			assert.Equal(t, tt.expected, got)
			if err != nil {
				assert.Equal(t, tt.errMsg, err.Error())
			}
		})
	}
}
