package main

import "fmt"

func getKeysGivenValue(mapHash map[string][]int, value int) []string {
	//
	keySlice := make([]string, 1)
	for key := range mapHash {
		for i := 0; i < len(mapHash[key]); i++ {
			if mapHash[key][i] == value {
				keySlice = append(keySlice, key)
			}
		}
	}
	return keySlice
}

func getMonthsofUniqueDays(dateMap map[string][]int) {
	//
}

func getDuplicateDay() {
	//
}
func getSingleValue() {
	//
}

func main() {
	dates := make(map[string][]int)
	dates["May"] = append(dates["May"], 15, 16, 19)
	dates["Jun"] = append(dates["Jun"], 17, 18)
	dates["Jul"] = append(dates["Jul"], 14, 16)
	dates["Aug"] = append(dates["Aug"], 14, 15, 17)

	for key := range dates {
		fmt.Println(key)
		for i := 0; i < len(dates[key]); i++ {
			fmt.Println(dates[key][i])
		}
	}

	fmt.Println(getKeysGivenValue(dates, 19))
	fmt.Println(getKeysGivenValue(dates, 14))

}
