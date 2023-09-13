package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
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

	url := "http://profiler:8081/create"

	requestBody := map[string]interface{}{
		"Login": id,
		"Email": wellFormatedEmail,
	}
	resp := postDataProfiler(url, requestBody)
	fmt.Println(resp) // FIXME

	if !AddUser(cred.Uri, id, psw, string(binary), wellFormatedEmail) {
		return "Unknown error while registration"
	}
	return "200"
}

func postProfileHandler(endpoint, userID, data string) string {
	if endpoint == "Description" && len(data) > 350 {
		return "Too long description"
	}
	if endpoint == "FullName" && len(data) > 30 {
		return "Too long Full Name"
	}
	if endpoint == "PhoneNB" && len(data) > 15 {
		return "Too long PhoneNB"
	}
	if endpoint == "Email" && len(data) > 50 {
		return "Too long Email"
	}
	if endpoint == "ProfilePic" && len(data) > 150 {
		return "Too long ProfilePic"
	}

	data = checkForBannedWords(data)

	requestBody := map[string]interface{}{
		"UserID": userID,
		"Data":   data,
	}

	url := "http://profiler:8081/" + endpoint
	postDataProfiler(url, requestBody)
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
	// requester, _ := zmq.NewSocket(zmq.PAIR)
	// defer requester.Close()
	// requester.Connect("tcp://profiler:5000")

	// for request_nbr := 0; request_nbr != 10; request_nbr++ {
	// 	// send hello
	// 	msg := fmt.Sprintf("Hello %d", request_nbr)
	// 	fmt.Println("Sending ", msg)
	// 	requester.Send(msg, 0)

	// 	// Wait for reply:
	// 	reply, _ := requester.Recv(0)
	// 	fmt.Println("Received ", reply)
	// }

	return true
}

func deleteUserData(userID string) string {
	requestBody := map[string]interface{}{
		"userID": userID,
	}

	resp := postDataProfiler("http://profiler:8081/delete", requestBody)
	return resp
}

func checkForBannedWords(inputString string) string {
	// Charger la liste des mots bannis depuis le fichier "lib.txt"
	bannedWords := loadBannedWords()

	// Diviser la chaîne d'entrée en mots individuels
	words := strings.Fields(inputString)

	// Vérifier si l'un des mots est dans la liste des mots bannis
	for i, word := range words {
		if contains(bannedWords, word) {
			// Si un mot banni est trouvé, le remplacer par des astérisques
			words[i] = strings.Repeat("*", len(word))
		}
	}

	// Reconstituer la chaîne de caractères à partir des mots modifiés
	return strings.Join(words, " ")
}

func loadBannedWords() []string {
	// Ouvrir le fichier "lib.txt" en lecture
	file, err := os.Open("lib.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		os.Exit(1)
	}
	defer file.Close()

	var bannedWords []string

	// Lire chaque ligne du fichier et ajouter les mots à la liste des mots bannis
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			bannedWords = append(bannedWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	return bannedWords
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
