package utils

import "strconv"

func ToUint(v string) uint {
	if v == "" {
		return 0
	}

	n, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0
	}

	return uint(n)
}
