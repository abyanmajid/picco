package utils

import (
	"fmt"
	"regexp"
)

func ExtractClassName(code string) (string, error) {
	re := regexp.MustCompile(`public\s+class\s+(\w+)`)
	matches := re.FindStringSubmatch(code)
	if len(matches) < 2 {
		return "", fmt.Errorf("could not find public class name")
	}
	return matches[1], nil
}
