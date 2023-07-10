package testcfg

import (
	"strconv"
)

func MaxTraces() []string {
	return GenerateTrace(3, 10)
}

func GenerateTrace(major, minor int) []string {
	res := make([]string, 0, major+minor)
	for i := 1; i <= major; i++ {
		res = append(res, "10"+strconv.Itoa(i))
	}
	for i := 1; i <= minor && i <= 9; i++ {
		res = append(res, "20"+strconv.Itoa(i))
	}
	for i := 10; i <= minor; i++ {
		res = append(res, "2"+strconv.Itoa(i))
	}
	return res
}
