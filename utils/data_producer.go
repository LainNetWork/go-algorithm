package utils

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//数据发生器，用于产生各种结构的数据样本

// GenerateArrayIntSorted
// @param n 生成的数量
func GenerateArrayIntSorted(n int) []int {
	arrayInt := GenerateArrayInt(n)
	sort.Ints(arrayInt)
	return arrayInt
}

// GenerateArrayInt 生成随机数组
// @param n 生成的数量
func GenerateArrayInt(n int) []int {
	intArray := make([]int, 0)
	for i := 0; i < n; i++ {
		intArray = append(intArray, rand.Intn(n))
	}
	return intArray
}
