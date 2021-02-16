package helper

func RemoveElementByValue(s []int, r int) []int {
	var filteredArray []int
	for _, v := range s {
		if v != r {
			filteredArray = append(filteredArray, v)
		}
	}
	return filteredArray
}

func GetUniqueValue(arr []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}

func FindIndexByValue(s []int, r int) int {
	var indexValue int
	for i, v := range s {
		if v == r {
			indexValue = i
			break
		}
	}
	return indexValue
}

func RemoveElementByValueLimit(s []int, r int, limit int) []int {
	var filteredArray []int
	var count = 0
	for _, v := range s {
		if v != r {
			filteredArray = append(filteredArray, v)
		} else {
			if count < limit {
				count++
			} else {
				filteredArray = append(filteredArray, v)
			}
		}
	}
	return filteredArray
}
