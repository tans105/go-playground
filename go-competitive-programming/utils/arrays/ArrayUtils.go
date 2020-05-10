package arrays

import (
	"strconv"
	"strings"
)

func ParseArray(str string, delimiter string) []int {
	str = strings.TrimSpace(str)
	effectiveStr := str[1:len(str)-1]
	elements := strings.Split(effectiveStr, delimiter)
	var intElements []int

	for _, i := range elements {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intElements = append(intElements, j)
	}

	return intElements
}

func Parse2DArray(str string, del byte) [][]int {
	str = strings.TrimSpace(str)
	str = str[1:len(str)-1]
	counter := 0
	start := false
	var entity []int
	var entities [][]int

	for counter != len(str) {
		if start && str[counter] != del && str[counter] != ']' {
			j, err := strconv.Atoi(string(str[counter]))
			if err != nil {
				panic(err)
			}
			entity = append(entity, j)
		}

		if str[counter] == '[' {
			start = true
		} else if str[counter] == ']' {
			entities = append(entities, entity)
			entity = nil
			start = false
		}

		counter++
	}

	return entities
}
