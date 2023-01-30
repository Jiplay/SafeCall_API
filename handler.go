package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
	zmq "github.com/pebbe/zmq4"
	"google.golang.org/protobuf/proto"
)

type Credentials struct {
	Uri         string `json:"uri"`
	AppPassword string `json:"appPassword"`
}

const (
	Login      int = 0
	Password   int = 1
	Reset      int = 2
	DeleteCode int = 3
)

func genUUID() string {
	id := uuid.New()
	return id.String()
}

// This function will get the uri in the json file to id to the db
func getCredentials() Credentials {
	fileContent, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
		return Credentials{}
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	res := Credentials{}
	json.Unmarshal([]byte(byteResult), &res)

	return res
}

func LoginHandler(id, psw string) string {
	cred := getCredentials()
	users := GetUsers(cred.Uri, "loginInfo")

	for _, info := range users {
		if info["login"] == id && info["psw"] == psw {
			userID := info["login"].(string)
			return userID
		}
	}

	return "failed"
}

func userToProto(username, psw string) UserMessage {
	user := UserMessage{
		Id:       1,
		Username: username,
		Password: psw,
		Settings: &SettingsMessage{
			DoNotDisturb: true,
			Language:     "eng",
		},
	}
	return user
}

func RegisterHandler(id, psw, email string) string { // TODO Ajouter un call au service de messagerie
	cred := getCredentials()
	users := GetUsers(cred.Uri, "loginInfo")

	re := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	wellFormatedEmail := re.FindString(email)

	nb, up, sp, le := verifyPassword(psw)

	if !nb {
		return "Your password must contains at least 1 digit"
	} else if !up {
		return "Your password must contains at least 1 uppercase"
	} else if !sp {
		return "Your password must contains at least 1 special character"
	} else if !le {
		return "Your password must contains at least 8 characters"
	}

	if len(id) < 5 {
		return "id too short" // Id too short
	} else if len(wellFormatedEmail) < 5 {
		return "Bad formated email"
	}

	for _, info := range users {
		if info["login"] == id {
			return "Id already taken" // Id already taken
		} else if info["email"] == wellFormatedEmail {
			return "email address already used"
		}
	}

	protoUser := userToProto(id, psw)
	binary, _ := proto.Marshal(&protoUser)
	if !CreateProfile(cred.Uri, id, wellFormatedEmail) {
		return "Unknown error while profile creation"
	}
	if !AddUser(cred.Uri, id, psw, string(binary), wellFormatedEmail) {
		return "Unknown error while registration"
	}
	return "200"
}

func postProfileHandler(endpoint, userID, data string) string {
	cred := getCredentials()
	if len(userID) > 45 {
		return "User ID too long"
	}

	if endpoint == "FullName" {
		if len(data) > 30 {
			return "Full Name too long"
		}
	}

	if endpoint == "PhoneNB" {
		if len(data) > 15 {
			return "Phone NB too long"
		}
	}

	UpdateProfile(cred.Uri, endpoint, userID, data)
	return "success"
}

func getProfileHandler(userID string) Profile {
	resp := getProfile(userID)

	return StringToProfile(resp)
}

func searchName(userID string) string {
	resp := searchNameQuery(userID)
	tmp := 1
	dest := strings.Split(resp, "\"")
	result := ""

	a := len(dest) - 2
	a = a / 4

	for a != 0 {
		tmp = tmp + 4
		result = result + "," + dest[tmp]
		a = a - 1
	}

	return result[1:]
}

func sendCallService(senter, dest string) bool {
	//  Socket to talk to server
	fmt.Println("Connecting to hello world server...")
	requester, _ := zmq.NewSocket(zmq.PAIR)
	defer requester.Close()
	requester.Connect("tcp://profiler:5555")

	for request_nbr := 0; request_nbr != 10; request_nbr++ {
		// send hello
		msg := fmt.Sprintf("Hello %d", request_nbr)
		fmt.Println("Sending ", msg)
		requester.Send(msg, 0)

		// Wait for reply:
		reply, _ := requester.Recv(0)
		fmt.Println("Received ", reply)
	}

	return true
}

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
		fmt.Println(info["email"], "David")
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
