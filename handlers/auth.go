package handlers

import (
	"crypto/subtle"
	"models" 
	"utils"
)

func Login(username, password string) (*models.User, error) {
	// TODO: implement database lookup
	hashedPassword := utils.HashPassword(password)
	if subtle.ConstantTimeCompare([]byte(hashedPassword), []byte("expected")) == 1 {
		return &models.User{Name: username}, nil
	}
	return nil, errors.New("invalid credentials")
}
