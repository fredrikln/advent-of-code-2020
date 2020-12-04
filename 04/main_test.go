package main

import (
	"reflect"
	"testing"

	"github.com/fredrikln/advent-of-code-2020/utils"
)

func Test_parseInput(t *testing.T) {
	testdata := utils.ReadInput("testinput.txt")

	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{testdata}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.data); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("parseInput() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	testdata := utils.ReadInput("testinput.txt")
	passports := parseInput(testdata)

	type args struct {
		passports []map[string]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{passports}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.passports); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	testdata := utils.ReadInput("testinput2.txt")
	passports := parseInput(testdata)

	type args struct {
		passports []map[string]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{passports}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.passports); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
