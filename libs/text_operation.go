package libs

import (
	"regexp"

	"github.com/jameskeane/bcrypt"
)

func HashPassword(pwd string) string {
	salt, err := bcrypt.Salt(10)
	if err != nil {
		return ""
	}
	hashedPW, err := bcrypt.Hash(pwd, salt)
	if err != nil {
		return ""
	}
	return hashedPW
}

func VerifyMailFormat(mail string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` 
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(mail)
}