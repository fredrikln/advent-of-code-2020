package main

import (
	"reflect"
	"testing"
)

func Test_parseGroups(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []group
	}{
		{
			"Example 1",
			args{[]string{"abcx", "abcy", "abcz"}},
			[]group{
				{3, map[string]int{"a": 3, "b": 3, "c": 3, "x": 1, "y": 1, "z": 1}},
			},
		},
		{
			"Example 2",
			args{[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			}},
			[]group{
				{1, map[string]int{"a": 1, "b": 1, "c": 1}},
				{3, map[string]int{"a": 1, "b": 1, "c": 1}},
				{2, map[string]int{"a": 2, "b": 1, "c": 1}},
				{4, map[string]int{"a": 4}},
				{1, map[string]int{"b": 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGroups(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGroups() = %v, want %v", got, tt.want)
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
		{
			"Example 1",
			args{[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			}},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{[]string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
