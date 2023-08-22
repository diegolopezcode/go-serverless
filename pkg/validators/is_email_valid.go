package validators

import "regexp"

func IsEmailValid(email string) bool {
	var rgxEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)

	if len(email) < 3 || len(email) > 254 || !rgxEmail.MatchString(email) {
		return false
	}
	return true
}
