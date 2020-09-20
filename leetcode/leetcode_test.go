package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 2, 3, 4, 5, 5, 5, 5, 5, 6, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteZero(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{{
		name: "test1",
		args: args{[]int{1, 2, 0, 3, 4, 5}},
		want: []int{1, 2, 3, 4, 5, 0},
	}, {
		name: "test2",
		args: args{[]int{1, 2, 0, 3, 0, 0}},
		want: []int{1, 2, 3, 0, 0, 0},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteZero(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
