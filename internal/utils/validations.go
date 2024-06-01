package utils

import (
	"fmt"
	"regexp"
)

func ParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
