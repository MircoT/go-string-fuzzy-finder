package alg

// Calculate the Levenshtein distance between target and the input.
func Levenshtein(target, input string) int { // nolint: funlen
	// Example
	// ref: https://www.datacamp.com/community/tutorials/fuzzy-string-python
	lenTarget := len(target)
	lenInput := len(input)
	numCols := lenTarget + 1

	distance := make([]int, (lenTarget+1)*(lenInput+1))

	// log.Println(target, input)

	for row := 0; row <= lenInput; row++ {
		for col := 0; col <= lenTarget; col++ {
			if row == 0 {
				distance[row*numCols+col] = col
			}

			if col == 0 {
				distance[row*numCols+col] = row
			}

			if row > 0 && col > 0 {
				cost := 0

				if target[col-1] != input[row-1] {
					cost++
				}

				// Costs
				// - x is the current posizion
				//
				// +-------------------+----------------+
				// | substitution cost | insetion cost  |
				// +-------------------+----------------+
				// | delete cost       |       x        |
				// +-------------------+----------------+
				delCost := distance[(row-1)*numCols+col] + 1        // Cost of deletions
				insCost := distance[row*numCols+(col-1)] + 1        // Cost of insertions
				subCost := distance[(row-1)*numCols+(col-1)] + cost // Cost of substitutions

				min := delCost

				if insCost < min {
					min = insCost
				}

				if subCost < min {
					min = subCost
				}

				distance[row*numCols+col] = min
			}
		}
	}

	// log.Printf("%+v", distance)

	// for row := 0; row <= lenInput; row++ {
	// 	for col := 0; col <= lenTarget; col++ {
	// 		fmt.Printf("%d ", distance[row*numCols+col])
	// 	}
	// 	fmt.Println()
	// }

	return distance[len(distance)-1]
}
