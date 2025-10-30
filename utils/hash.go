package utils

import "golang.org/x/crypto/bcrypt"

/*
*
Why use 14 in the GenerateFromPassword function?

The cost factor determines how computationally expensive the hashing operation is.

A higher cost means more secure, but also slower to compute.

Common values:

10 -> reasonable speed and security for most applications

14 -> stronger but slower (good for servers with more resources)

Essentially, bcrypt will run more internal rounds to make brute-force attacks harder.
*
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Check the has for the password and the hashedPassword match
func CheckPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
