package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]int{366, 675, 1456, 979, 299, 1721}}, 514579},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.expenses); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]int{299, 1456, 1721, 979, 366, 675}}, 241861950},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.expenses); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
