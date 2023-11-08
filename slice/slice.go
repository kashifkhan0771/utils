package slice

import (
	"fmt"

	"github.com/kashifkhan0771/utils/maps"
)

// RemoveDuplicateStr removes all the duplicate strings and return a new slice without any duplicate values.
func RemoveDuplicateStr(strSlice []string) []string {
	duplicate := maps.NewStateMap()

	newSlice := make([]string, 0)

	for _, value := range strSlice {
		if duplicate.HasState(value) {
			continue
		}

		duplicate.SetState(value, true)
		newSlice = append(newSlice, value)
	}

	return newSlice
}

// RemoveDuplicateInt removes all the duplicate integers and return a new slice without any duplicate values.
func RemoveDuplicateInt(strSlice []int) []int {
	duplicate := maps.NewStateMap()

	newSlice := make([]int, 0)

	for _, value := range strSlice {
		if duplicate.HasState(fmt.Sprintf("%d", value)) {
			continue
		}

		duplicate.SetState(fmt.Sprintf("%d", value), true)
		newSlice = append(newSlice, value)
	}

	return newSlice
}
