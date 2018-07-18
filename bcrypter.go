package main

import "os"
import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {
	password := os.Args[1]
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
