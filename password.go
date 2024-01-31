package main

import (
	"fmt"
	"math/rand"
)

type Login struct {
	Website  string
	Username string
	Password string
}

func main() {
	meme := Login{"website", "username", generatePassword(12)}
	fmt.Println("login1", meme)
	meme.UpdateLogin()
	fmt.Println("new login", meme)
}

func generatePassword(length int) string {
	var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "!", "@", "&", "?"}
	password := ""
	for i := 0; i < length; i++ {
		r := rand.Intn(len(letters))
		letter := letters[r]
		password += letter
	}
	return password
}

func (l *Login) GenerateLogin() {
	l.Website = "meme.com"
	l.Username = "xx42069xx@meme.com"
	l.Password = generatePassword(5)
}

func (l *Login) UpdateLogin() {
	fmt.Print("New Username: ")
	fmt.Scan(&l.Username)
	fmt.Print("New Password (or leave blank for random): ")
	newPW, _ := fmt.Scanln(&l.Password)
	fmt.Print(newPW)
	if newPW == 0 {
		l.Password = generatePassword(12)
	}
}
