package models

import "testing"

func TestGenerateHashPassword(t *testing.T) {
	// Arrange
	// var password string = "Aa1029!@#$%^&*()-=_+"
	var password string = "A"
	var user *User = new(User)
	user.Password = password

	// Act
	var error = user.GenerateHashPassword()

	// Assert
	if error != nil {
		t.Error(error.Error())
	}

	if user.Password == password {
		t.Error("Хеш совпадает с исходным паролем.")
	}
}
