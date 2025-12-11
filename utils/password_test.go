package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {

	//create pass random
	password := RandomString(8)

	//hash Password
	hashPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)

	err = CheckPassword(password, hashPassword1)
	require.NoError(t, err)

	wrongPassword := RandomString(8)
	err = CheckPassword(wrongPassword, hashPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	//hash Password
	hashPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)
	require.NotEqual(t, hashPassword1, hashPassword2)
}
