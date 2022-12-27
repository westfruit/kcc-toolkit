package convert

import (
	"strconv"
	"strings"
)

// 合并
func Merge[T int | int8 | int32 | int64 | string](arrays ...[]T) []T {
	var result []T
	for _, array := range arrays {
		result = append(result, array...)
	}

	return result
}

// 合并且去重
func MergeAndRemoveDuplicate[T int | int8 | int32 | int64 | string](arrays ...[]T) []T {
	var result []T
	if len(arrays) == 0 {
		return result
	}

	for _, array := range arrays {
		result = append(result, array...)
	}

	return RemoveDuplicate(result)
}

// 去重
func RemoveDuplicate[T int | int8 | int32 | int64 | string](arr []T) []T {
	var result []T
	if len(arr) == 0 {
		return result
	}

	tempMap := map[T]byte{}

	for _, e := range arr {
		if _, ok := tempMap[e]; !ok {
			tempMap[e] = 1
			result = append(result, e)
		}
	}

	return result
}

// 查找元素s在array中是否存在
func IndexOf(array []string, s string) int {
	for index, value := range array {
		if value == s {
			return index
		}
	}

	return -1
}

// 字符串数组转成int64数组
func StringArrayToInt64Array(sa []string) ([]int64, error) {
	si := make([]int64, 0, len(sa))
	for _, a := range sa {
		n, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return si, err
		}
		si = append(si, n)
	}
	return si, nil
}

// 字符串转成int64数组
func StringToInt64Array(str string) ([]int64, error) {
	str = strings.Trim(str, ",")
	ns := strings.Split(str, ",")

	return StringArrayToInt64Array(ns)
}

// Contains 数组是否包含某元素
func Contains[T int | int8 | int32 | int64 | string](slice []T, s T) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}

	return -1
}

// 字符串是否包含其它字符串
func ContainsSubString(src string, sub string) bool {
	if len(src) == 0 {
		return false
	}

	arr := strings.Split(src, ",")
	n := Contains(arr, sub)

	return n >= 0
}
