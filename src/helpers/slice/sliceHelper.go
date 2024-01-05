//go:generate generate-interfaces.sh

package helpers

// CombineOnlyMatchingValues returns a new slice that only contains values that are in both slices.
func CombineOnlyMatchingValues(sliceOne, sliceTwo []string) (combination []string) {
	for sliceOneIndex := range sliceOne {
		valueOne := sliceOne[sliceOneIndex]

		for sliceTwoIndex := range sliceTwo {
			valueTwo := sliceTwo[sliceTwoIndex]

			if valueOne == valueTwo {
				combination = append(combination, valueOne)
				break
			}
		}
	}

	return combination
}
