package main

import "fmt"

func position(slice []int, value int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

func getKeysGivenValue(mapHash map[string][]int, value int) []string {
	var keySlice []string
	for key := range mapHash {
		for i := 0; i < len(mapHash[key]); i++ {
			if mapHash[key][i] == value {
				keySlice = append(keySlice, key)
			}
		}
	}
	return keySlice
}

func getMonthsofUniqueDays(dateMap map[string][]int) []string {
	var months []string
	for day := 14; day < 20; day++ {
		tempMonth := getKeysGivenValue(dateMap, day)
		if len(tempMonth) == 1 {
			months = append(months, tempMonth[0])
		}
	}
	return months
}

func getDuplicateDay(dateMap map[string][]int) int {
	for day := 14; day < 20; day++ {
		tempMonth := getKeysGivenValue(dateMap, day)
		if len(tempMonth) == 2 {
			return day
		}
	}
	return 0
}

func removeFromSlice(slice []int, pos int) []int {
	return slice[:pos+copy(slice[pos:], slice[pos+1:])]
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

	//First statement
	fmt.Println("Starting dates: ", dates)
	fmt.Println("Cheryl then tells Albert and Bernard the month and the day of her birthday, respectively.")
	fmt.Println("Albert: I don't know when Cheryl's birthday is, but I know that Bernard doesn't know either.")
	uniqueMonths := getMonthsofUniqueDays(dates)
	for i := 0; i < len(uniqueMonths); i++ {
		delete(dates, uniqueMonths[i])
	}
	fmt.Println("    Dates after Albert's statement: ", dates)

	//Second statement
	duplicateDay := getDuplicateDay(dates)
	i := position(dates["Jul"], duplicateDay)
	dates["Jul"] = dates["Jul"][:i+copy(dates["Jul"][i:], dates["Jul"][i+1:])]
	i = position(dates["Aug"], duplicateDay)
	dates["Aug"] = dates["Aug"][:i+copy(dates["Aug"][i:], dates["Aug"][i+1:])]
	fmt.Println("Bernard: At first I didn't know when Cheryl's birthday is, but now I know.")
	fmt.Println("    Dates after Bernard's statement: ", dates)

	//Third statement
	for key := range dates {
		if len(dates[key]) > 1 {
			delete(dates, key)
		}
	}
	fmt.Println("Albert: Now I know when Cheryl's birthday is too.")
	fmt.Println("    Date after Albert's final statement: ", dates)
}
