package utils

import (
	"fmt"
	"strings"
)

// formatCurrency returns the given number, formated as currency
func FormatCurrency(n float64) string {
	s := fmt.Sprintf("%.2f", n)

	// pisahkan integer dan desimal
	parts := strings.Split(s, ".")
	intPart := parts[0]
	decPart := parts[1]

	// kasih koma ribuan di intPart
	var result []string
	for i, c := range intPart {
		if (len(intPart)-i)%3 == 0 && i != 0 {
			result = append(result, ",")
		}
		result = append(result, string(c))
	}

	return strings.Join(result, "") + "." + decPart
}
