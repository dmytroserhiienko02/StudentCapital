package main

import (
	"reflect"
	"testing"
)

func TestGetMaxCapital(t *testing.T) {
	laptops1 := []Laptop{
		{Gain: 3, Price: 4},
		{Gain: 2, Price: 5},
		{Gain: 1, Price: 6},
		{Gain: 5, Price: 4},
		{Gain: 7, Price: 3},
		{Gain: 7, Price: 4},
		{Gain: 8, Price: 2},
		{Gain: 12, Price: 6},
		{Gain: 14, Price: 8},
	}
	laptops2 := []Laptop{
		{Gain: 20, Price: 5},
		{Gain: 5, Price: 2},
		{Gain: 2, Price: 1},
	}
	laptops3 := []Laptop{
		{Gain: 5, Price: 20},
		{Gain: 2, Price: 5},
		{Gain: 1, Price: 2},
	}
	laptops4 := []Laptop{
		{Gain: 10, Price: 5},
	}
	laptops5 := []Laptop{
		{Gain: 5, Price: 10},
	}

	type args struct {
		N       int
		C       float64
		laptops []Laptop
	}
	tests := []struct {
		name   string
		args   args
		expect float64
	}{
		{
			name:   "General test",
			args:   args{N: 3, C: 5.0, laptops: laptops1},
			expect: 21.0,
		},
		{
			name:   "Test for 0 value ",
			args:   args{N: 0, C: 5.0, laptops: laptops1},
			expect: 5.0,
		},
		{
			name:   "Check iteration",
			args:   args{N: 3, C: 1.0, laptops: laptops2},
			expect: 20.0,
		},
		{
			name:   "Check iteration 2",
			args:   args{N: 2, C: 1.0, laptops: laptops2},
			expect: 5.0,
		},
		{
			name:   "Swapped price and gain",
			args:   args{N: 5, C: 1.0, laptops: laptops3},
			expect: 1.0,
		},
		{
			name:   "1 element",
			args:   args{N: 5, C: 5.0, laptops: laptops4},
			expect: 10.0,
		},
		{
			name:   "1 swapped element",
			args:   args{N: 5, C: 5.0, laptops: laptops5},
			expect: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxCapital(tt.args.N, tt.args.C, tt.args.laptops); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("getMaxCapital() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
