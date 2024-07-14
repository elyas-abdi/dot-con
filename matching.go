package con

import (
	`regexp`
	`strings`
)

func matchesAssignmentPattern(str string) bool {
	return regexp.MustCompile(`^\w+\s*=\s*.+$`).MatchString(str)
}

func matchesFactorsPattern(str string) bool {
	return strings.HasPrefix(str, factorsPrefix) && strings.HasSuffix(str, factorsSuffix)
}
