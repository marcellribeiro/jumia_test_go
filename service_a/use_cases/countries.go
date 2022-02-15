package use_cases

import (
	"regexp"
)

// GetCountryByPhone was created using the simple switch case because the given regex doesn't match with the phone list in CSV file
func GetCountryByPhone(phone string) string {
	phone = getOnlyNumbers(phone)
	phone = phone[0:3]

	switch phone {
	case "237":
		return "Cameroon"
	case "251":
		return "Ethiopia"
	case "212":
		return "Morocco"
	case "258":
		return "Mozambique"
	case "256":
		return "Uganda"
	default:
		return ""
	}
}

func getOnlyNumbers(phone string) string {
	pattern := regexp.MustCompile(`(\d+)`)
	numberStrings := pattern.FindAllStringSubmatch(phone, -1)
	numbers := ""
	for _, numberString := range numberStrings {
		number := numberString[1]
		numbers = numbers + number
	}
	return numbers
}
