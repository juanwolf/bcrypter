package main

import "os"
import "fmt"
import "strconv"
import "golang.org/x/crypto/bcrypt"

func main() {
	password := os.Args[1]
	cost := bcrypt.DefaultCost
	var err error
	if len(os.Args) > 2 {
		cost, err = strconv.Atoi(os.Args[2])
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
