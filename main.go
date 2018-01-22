// bitSort project main.go
package main

import (
	"fmt"
)

func main() {
	r := bitIn([]int32{1, 2, 3, 5, 8, 10, 1922, 123, 52356, 1345, 2342, 342, 34, 23, 42, 34, 234, 234}, 52356)
	fmt.Println(r)
	arr := bitOut(r)
	fmt.Println(arr)
}

// 位 排序
func bitIn(arr []int32, maxValue int32) (result []byte) {
	n := maxValue/8 + 1
	result = make([]byte, n)
	for _, v := range arr {
		if v > maxValue {
			fmt.Printf("err max=", maxValue, " now=", v)
			continue
		}
		result[int(v)/8] = getByte(result[int(v)/8], byte(v%8))
	}

	return
}

func bitOut(result []byte) (arr []int32) {
	n := len(result) * 8
	for i := 0; i < n; i++ {
		//		fmt.Println(i/8, i%8, getByteToIndex(result[i/8], i%8))
		v := getByteToInt32(getByteToIndex(result[i/8], i%8))
		if v != 3 {
			v += int32(i/8) * 8
			arr = append(arr, v)
		}
	}
	return
}

func getByte(value, add byte) byte {
	switch add {
	case 0:
		value |= 1
	case 1:
		value |= 2
	case 2:
		value |= 4
	case 3:
		value |= 8
	case 4:
		value |= 16
	case 5:
		value |= 32
	case 6:
		value |= 64
	case 7:
		value |= 128
	}
	return value
}

func getByteToIndex(value byte, index int) byte {
	switch index {
	case 0:
		return value & 1
	case 1:
		return value & 2
	case 2:
		return value & 4
	case 3:
		return value & 8
	case 4:
		return value & 16
	case 5:
		return value & 32
	case 6:
		return value & 64
	case 7:
		return value & 128
	}
	return 3
}

func getByteToInt32(value byte) int32 {
	switch value {
	case 1:
		return 0
	case 2:
		return 1
	case 4:
		return 2
	case 8:
		return 3
	case 16:
		return 4
	case 32:
		return 5
	case 64:
		return 6
	case 128:
		return 7
	}
	return 3
}
