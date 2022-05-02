package kata

import "strings"

func toWeirdCase(str string) string {
	count := 0
	str1 := strings.Split(str, "")
	strRet := ""
	for i, _ := range str1 {
		if str1[i] != " " {
			if count%2 == 0 {
				str1[i] = strings.ToUpper(str1[i])
			} else {
				str1[i] = strings.ToLower(str1[i])
			}
			count += 1
		} else {
			count = 0
		}

		strRet = strRet + str1[i]
	}
	return strRet
}
