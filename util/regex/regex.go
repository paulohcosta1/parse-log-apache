package regex

func GetValueByGroup(nameGroup string, match []string, groupNames []string) string {

	for groupIdx, value := range match {
		name := groupNames[groupIdx]
		if name == nameGroup {
			return value
		}
	}
	return ""
}
