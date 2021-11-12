package crypto_utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hashedPwd string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) == nil
}
