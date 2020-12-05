package main

import (
	"testing"
)

func Test_convertRow(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"BFFFBBFRRR", args{"BFFFBBFRRR"}, 70},
		{"FFFBBBFRRR", args{"FFFBBBFRRR"}, 14},
		{"BBFFBBFRLL", args{"BBFFBBFRLL"}, 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertRow(tt.args.row); got != tt.want {
				t.Errorf("convertRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertColumn(t *testing.T) {
	type args struct {
		column string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"BFFFBBFRRR", args{"BFFFBBFRRR"}, 7},
		{"FFFBBBFRRR", args{"FFFBBBFRRR"}, 7},
		{"BBFFBBFRLL", args{"BBFFBBFRLL"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertColumn(tt.args.column); got != tt.want {
				t.Errorf("convertColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
