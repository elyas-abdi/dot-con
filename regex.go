package con

import `regexp`

func matchesAssignment(str string) bool {
	return regexp.MustCompile(`^\w+\s*=\s*.+$`).MatchString(str)
}
