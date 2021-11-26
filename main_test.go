package main

import (
	"reflect"
	"testing"
)

func TestInstancesOfAnd(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Test cases
			args: args{
				text: "Maisy and nell and finn",
			},
			want: 2,
		},
		{
			args: args{
				text: "test",
			},
			want: 0,
		},
		{
			args: args{
				text: "andandand",
			},
			want: 0,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InstancesOfAnd(tt.args.text); got != tt.want {
				t.Errorf("InstancesOfAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOutputResponse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			// Test cases
			args: args{
				text: "",
			},
			want: Response{"Invalid String", "N/A", "404"},
		},
		{
			args: args{
				text: "Hi there",
			},
			want: Response{"Hi there", "0", "200"},
		},
		{
			args: args{
				text: "andandand",
			},
			want: Response{"andandand", "0", "200"}, //only counts instances of and used as a word rather than the 3 characters repeated
		}, {
			args: args{
				text: "Maisy and Nell and Finn",
			},
			want: Response{"Maisy and Nell and Finn", "2", "200"},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OutputResponse(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OutputResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
