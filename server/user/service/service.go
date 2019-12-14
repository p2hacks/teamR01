package service

import (
    "golang.org/x/crypto/bcrypt"
)

func CreateUserId(str string) (string, error){
    hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}