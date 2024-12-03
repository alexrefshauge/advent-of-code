package parsing

import "strings"

func strToRows(in string) []string {
	return strings.Split(in, "\n")
}
