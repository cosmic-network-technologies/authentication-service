package main

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(*password), 8)

	return hash
}

func VerifyPassword(hash *[]byte, password *string) bool {
	return bcrypt.CompareHashAndPassword(*hash, []byte(*password)) == nil
}
