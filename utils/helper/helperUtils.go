package helper

// ContainsString check if given element to search exists in the given string array
func ContainsString(array []string, searchElem string) bool {
	for _, elem := range array {
		if elem == searchElem {
			return true
		}
	}
	return false
}

// AppendIfMissingString append ints to an existing int array if it doesn't contain it already
func AppendIfMissingString(array []string, elemsToAdd ...string) []string {
	for _, elemToAdd := range elemsToAdd {
		alreadyExist := false
		for _, elem := range array {
			if elem == elemToAdd {
				alreadyExist = true
				break
			}
		}
		if !alreadyExist {
			array = append(array, elemToAdd)
		}
	}
	return array
}
