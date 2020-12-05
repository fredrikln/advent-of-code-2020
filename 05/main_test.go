package main

import (
	"testing"
)

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

func Test_getID(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"BFFFBBFRRR", args{"BFFFBBFRRR"}, 567},
		{"FFFBBBFRRR", args{"FFFBBBFRRR"}, 119},
		{"BBFFBBFRLL", args{"BBFFBBFRLL"}, 820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getID(tt.args.row); got != tt.want {
				t.Errorf("getID() = %v, want %v", got, tt.want)
			}
		})
	}
}
