package stmt

import "fmt"

func BuildWhere(whereValuesMap map[string][]string) string {

	if len(whereValuesMap) == 0 {
		return ""
	}

	whereCondition := " WHERE "
	mapIndex := 1

	for key, array := range whereValuesMap {

		for index, value := range array {

			whereCondition += fmt.Sprintf("%s = '%s'", key, value)
			if index != len(array)-1 {
				whereCondition += "OR "
			}
		}

		if mapIndex != len(whereValuesMap) {
			whereCondition += "AND "
		}
		mapIndex++
	}
	return whereCondition
}
