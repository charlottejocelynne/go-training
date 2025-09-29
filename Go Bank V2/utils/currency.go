package utils

import (
	"fmt"
	"strings"
)

func FormatCurrency(n float64) string {
	s := fmt.Sprintf("%.2f", n)
	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := parts[1]

	var result []string
	for i, c := range intPart {
		if (len(intPart)-i)%3 == 0 && i != 0 {
			result = append(result, ",")
		}
		result = append(result, string(c))
	}
	return strings.Join(result, "") + "." + decPart
}
