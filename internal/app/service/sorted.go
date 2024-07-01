package service

import (
	"strconv"
)

func ParamsToArr(params string) []int {
	var num string
	var counter int
	numArr := make([]int, arrSize(params))

	for i := 0; i < len(params); i++ {

		if params[i] == ',' {
			intVar, _ := strconv.Atoi(num)
			numArr[counter] = intVar
			counter++
			num = ""
		} else {
			num += string(params[i])

			if i == len(params)-1 {
				intVar, _ := strconv.Atoi(num)
				numArr[counter] = intVar
				counter++
			}
		}

	}

	sortedArr(numArr)
	return numArr

}

func arrSize(params string) int {
	counter := 1
	for i := 0; i < len(params); i++ {
		if params[i] == ',' && i != len(params)-1 {
			counter++
		}
	}
	return counter
}

func sortedArr(numArr []int) {

	for i := 0; i < len(numArr); i++ {
		for j := i + 1; j < len(numArr); j++ {

			if numArr[i] > numArr[j] {
				numArr[i], numArr[j] = numArr[j], numArr[i]
			}
		}
	}
}
