package main

import "testing"

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
