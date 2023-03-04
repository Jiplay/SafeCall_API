package main

import "unicode"

func PasswordHandler(login, old, new string) string {
	cred := getCredentials()
	users := GetUsers(cred.Uri, "loginInfo")

	nb, up, sp, le := verifyPassword(new)

	if !nb {
		return "Your password must contains at least 1 digit"
	} else if !up {
		return "Your password must contains at least 1 uppercase"
	} else if !sp {
		return "Your password must contains at least 1 special character"
	} else if !le {
		return "Your password must contains at least 8 characters"
	}

	for _, info := range users {
		if info["login"] == login && info["psw"] == old {
			err := editLoginInfo(cred.Uri, login, new, Password)
			if !err {
				return "network error"
			}
			return "200"
		}
	}
	return "user not found"
}

func verifyPassword(s string) (number, upper, special, length bool) {
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		default:
			//return false, false, false, false
		}
	}
	length = len(s) >= 7

	return
}

func ForgetPasswordHandler(email string) string {
	cred := getCredentials()
	users := GetUsers(cred.Uri, "loginInfo")

	for _, info := range users {
		if info["email"] == email {
			code := codeGenerator()
			editLoginInfo(cred.Uri, email, code, Login)
			sendMail(cred.AppPassword, email, code)
			return "200"
		}
	}
	return "ko"
}

func checkCodeHandler(email, code string) bool {
	cred := getCredentials()
	users := GetUsers(cred.Uri, "loginInfo")

	for _, info := range users {
		if info["email"] == email && info["code"] == code {
			return true
		}
	}
	return false
}

func setPasswordHandler(email, password string) string {
	cred := getCredentials()
	nb, up, sp, le := verifyPassword(password)

	if !nb {
		return "Your password must contains at least 1 digit"
	} else if !up {
		return "Your password must contains at least 1 uppercase"
	} else if !sp {
		return "Your password must contains at least 1 special character"
	} else if !le {
		return "Your password must contains at least 8 characters"
	}

	err := editLoginInfo(cred.Uri, email, password, Reset)

	if !err {
		return "error"
	}
	err = editLoginInfo(cred.Uri, email, "", Login)
	if !err {
		return "error"
	}
	return "200"
}
