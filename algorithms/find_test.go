package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		origin []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "奇数数量", args: args{origin: []int{1, 2, 3, 4, 5}, target: 4}, want: 3},
		{name: "偶数数量", args: args{origin: []int{1, 2, 3, 4}, target: 2}, want: 1},
		{name: "没有找到", args: args{origin: []int{1, 2, 3, 4}, target: 42}, want: -1},
		{name: "没有找到", args: args{origin: []int{1, 2, 3, 4}, target: -42}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.origin, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
