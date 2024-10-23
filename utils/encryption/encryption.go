package encryption

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(passHash),nil
}

func CheckPassword(password, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass),[]byte(password))
	return err == nil
}