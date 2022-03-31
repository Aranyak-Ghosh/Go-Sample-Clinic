package hash

import "golang.org/x/crypto/bcrypt"

func HashAndSaltPassword(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	} else {
		return hash, nil
	}
}

func VerifyPassword(pwd []byte, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pwd)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
