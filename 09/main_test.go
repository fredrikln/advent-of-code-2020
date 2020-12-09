package main

import (
	"reflect"
	"testing"
)

func Test_isValidNumber(t *testing.T) {
	type args struct {
		previous []int
		number   int
		preamble int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example 1",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				26,
				25,
			},
			true,
		},
		{
			"Example 2",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				49,
				25,
			},
			true,
		},
		{
			"Example 3",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				100,
				25,
			},
			false,
		},
		{
			"Example 4",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				50,
				25,
			},
			false,
		},
		{
			"Example 5",
			args{
				[]int{20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 45},
				26,
				25,
			},
			true,
		},
		{
			"Example 6",
			args{
				[]int{20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 45},
				65,
				25,
			},
			false,
		},
		{
			"Example 7",
			args{
				[]int{20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 45},
				64,
				25,
			},
			true,
		},
		{
			"Example 8",
			args{
				[]int{20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25, 45},
				66,
				25,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidNumber(tt.args.previous, tt.args.number, tt.args.preamble); got != tt.want {
				t.Errorf("isValidNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstInvalid(t *testing.T) {
	type args struct {
		numbers  []int
		preamble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 5}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstInvalid(tt.args.numbers, tt.args.preamble); got != tt.want {
				t.Errorf("firstInvalid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findContiguous(t *testing.T) {
	type args struct {
		numbers []int
		number  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127}, []int{15, 25, 47, 40}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContiguous(tt.args.numbers, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContiguous() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMax(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"Example 1", args{[]int{1, 2, 3, 4, 5}}, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := minMax(tt.args.numbers)
			if got != tt.want {
				t.Errorf("minMax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("minMax() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		numbers []int
		number  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127}, 62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.numbers, tt.args.number); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
