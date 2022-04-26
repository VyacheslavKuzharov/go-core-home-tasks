package main

import (
	"home-task9/users"
	"testing"
)

func Test_maxAge(t *testing.T) {
	type args[P users.User] struct {
		peoples []P
	}

	tests1 := []struct {
		name string
		args args[users.Employee]
		want int
	}{
		{
			name: "Employee test1",
			args: args[users.Employee]{
				peoples: []users.Employee{{CurrentAge: 22}, {CurrentAge: 8}},
			},
			want: 22,
		},
		{
			name: "Employee test2",
			args: args[users.Employee]{
				peoples: []users.Employee{{CurrentAge: -3}, {CurrentAge: 8}},
			},
			want: 8,
		},
		{
			name: "Employee test3",
			args: args[users.Employee]{
				peoples: []users.Employee{},
			},
			want: 0,
		},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAge[users.Employee](tt.args.peoples...); got != tt.want {
				t.Errorf("maxAge() = %v, want %v", got, tt.want)
			}
		})
	}

	tests2 := []struct {
		name string
		args args[users.Customer]
		want int
	}{
		{
			name: "Customer test1",
			args: args[users.Customer]{
				peoples: []users.Customer{{CurrentAge: 22}, {CurrentAge: 8}},
			},
			want: 22,
		},
		{
			name: "Customer test2",
			args: args[users.Customer]{
				peoples: []users.Customer{{CurrentAge: -3}, {CurrentAge: 8}},
			},
			want: 8,
		},
		{
			name: "Customer test3",
			args: args[users.Customer]{
				peoples: []users.Customer{},
			},
			want: 0,
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAge[users.Customer](tt.args.peoples...); got != tt.want {
				t.Errorf("maxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
