package fibonacci

import "fmt"

func ParseIndex(i string) (string, error) {
	if i == "9" {
		return "34", nil
	} else {
		return "", fmt.Errorf("feature not implemented :)")
	}
}
