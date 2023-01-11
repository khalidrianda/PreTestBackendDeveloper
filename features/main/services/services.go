package services

import "testcase/features/main/entity"

type services struct {
	srv entity.Services
}

func New(srv entity.Services) entity.Services {
	return &services{srv: srv}
}

// 1. Buatkan fungsi yang bisa menentukan argumen input text adalah palindrom atau bukan.
func (as *services) Palindrome(n string) string {
	var temp string

	for i := len(n) - 1; i >= 0; i-- {
		temp += string(n[i])
	}

	for i := len(n) - 1; i >= 0; i-- {
		if n[i] != temp[i] {
			return "Not palindrome"
		}
	}

	return "Palindrome"
}
