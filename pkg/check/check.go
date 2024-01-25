package check

import (
 "errors"
 "unicode"
)

func Password(password string) error {
 if len(password) < 8 {
  return errors.New("error in password format")
 }

 hasUppercase := false
 hasLowercase := false
 hasDigit := false

 for _, char := range password {
  switch {
  case unicode.IsUpper(char):
   hasUppercase = true
  case unicode.IsLower(char):
   hasLowercase = true
  case unicode.IsDigit(char):
   hasDigit = true
  case unicode.IsSpace(char):
   return errors.New("password is wrong")
  }
 }

 if !hasUppercase || !hasLowercase || !hasDigit {
  return errors.New("password must contain uppercase, lowercase, and digit characters")
 }

 return nil
}
