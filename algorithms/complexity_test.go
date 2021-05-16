package algorithms

import (
	"fmt"
	"github.com/LainNetWork/go-algorithm/utils"
	"reflect"
	"testing"
)

//时间复杂度
//一个有序数组A，长度为M，另一个无序数组B,长度为N，打印B中所有不在A中的数，A数组长度为N，B数组长度为M。

//方法一：遍历所有B数组中的数，将B数组中的数与A数组一一比对
func TestComplexity_1(t *testing.T) {
	for i := 0; i < 500000; i++ {
		A := utils.GenerateArrayIntSorted(15)
		B := utils.GenerateArrayInt(10)
		complexity1 := Complexity1(A, B)
		fmt.Println(A)
		fmt.Println(B)
		fmt.Println(complexity1)
		if !checkComplexity(A, complexity1) {
			t.Error("结果不符合！")
		}
	}
}

//判断结果是否正确 (A中确实不包含结果数组元素)
func checkComplexity(A []int, result []int) bool {
	for _, r := range result {
		for _, o := range A {
			if o == r {
				return false
			}
		}
	}
	return true
}

func TestComplexity2(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "奇数数量", args: args{A: []int{1, 2, 3, 4, 5}, B: []int{4, 5, 6, 7, 8}}, want: []int{6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complexity2(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complexity2() = %v, want %v", got, tt.want)
			}
		})
	}
}
