package main

import "testing"

const exampleInput1 string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput1,
			want:  36,
		},
		{
			name:  "real",
			input: input,
			want:  674,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

const exampleInput2 string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  any
	}{
		{
			name:  "example",
			input: exampleInput2,
			want:  81,
		},
		{
			name:  "real",
			input: input,
			want:  1372,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
