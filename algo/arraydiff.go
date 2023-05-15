package main

//Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.
//
//It should remove all values from list a, which are present in list b keeping their order.
//
//array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
//If a value is present in b, all of its occurrences must be removed from the other:

func ArrayDiff(a, b []int) []int {
	result := []int{}

	for _, num := range a {
		isPresent := false
		for _, n := range b {
			if num == n {
				isPresent = true
				break
			}
		}
		if !isPresent {
			result = append(result, num)
		}
	}

	return result
}
